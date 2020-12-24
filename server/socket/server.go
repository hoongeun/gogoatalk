package socket

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"

	"github.com/Hoongeun/gogoatalk/common"
	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"github.com/Hoongeun/gogoatalk/server/core"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	jwtPrivateKey            *rsa.PrivateKey
	jwtPublicKey             *rsa.PublicKey
	am                       *core.AccountManager
	cm                       *core.ChatRoomManager
	gs                       *grpc.Server
	broadcaster              chan pb.StreamResponse
	sessionChans             map[string]chan pb.StreamResponse
	sessionMtx, broadcastMtx sync.RWMutex
	pb.UnimplementedRouteChatServer
}

func NewServer(jwtKeyPath string) (*Server, error) {
	pubKey, err := ioutil.ReadFile(jwtKeyPath + ".pub")
	if err != nil {
		return nil, fmt.Errorf("Error reading the jwt privatekey: %v", err)
	}
	privKey, err := ioutil.ReadFile(jwtKeyPath)
	if err != nil {
		return nil, fmt.Errorf("Error reading the jwt publickey: %v", err)
	}

	privatekey, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		return nil, fmt.Errorf("Error parsing the jwt privatekey: %s", err)
	}

	publickey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return nil, fmt.Errorf("Error parsing the jwt publickey: %s", err)
	}

	am := core.NewAccountManager()
	if err := am.LoadAccounts(); err != nil {
		log.Fatalf("Cannot Load Account: %s", err)
		return nil, err
	}
	cm := core.NewChatRoomManager()

	return &Server{
		jwtPrivateKey: privatekey,
		jwtPublicKey:  publickey,
		am:            am,
		cm:            cm,
		broadcaster:   make(chan pb.StreamResponse, 1000),
		sessionChans:  make(map[string]chan pb.StreamResponse),
	}, nil
}

func (s *Server) Listen(host string, ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s.gs = grpc.NewServer()
	pb.RegisterRouteChatServer(s.gs, s)

	l, err := net.Listen("tcp", host)
	if err != nil {
		return errors.New("server unable to bind on provided host")
	}

	fmt.Println(fmt.Sprintf("Server is Listening on %s", host))

	go s.broadcast(ctx)

	go func() {
		s.gs.Serve(l)
		cancel()
	}()

	<-ctx.Done()

	s.broadcaster <- pb.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &pb.StreamResponse_ServerShutdown{
			&pb.StreamResponse_Shutdown{},
		},
	}
	close(s.broadcaster)
	s.gs.GracefulStop()
	return nil
}

func (s *Server) Stream(srv pb.RouteChat_StreamServer) error {
	md, ok := metadata.FromIncomingContext(srv.Context())
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	jwtToken, ok := md["authorization"]
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	if len(jwtToken) == 0 || jwtToken[0] == "" {
		return grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}
	go s.sendBroadcasts(srv, jwtToken[0])

	for {
		_, err := srv.Recv() // TODO
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		// TODO demux bi-streaming
		// s.SendMessage(srv.Context(), req)
	}

	<-srv.Context().Done()
	return srv.Context().Err()
}

func (s *Server) broadcast(ctx context.Context) {
	for res := range s.broadcaster {
		s.broadcastMtx.RLock()
		for _, c := range s.sessionChans {
			select {
			case c <- res:
				// noop
			default:
				common.ServerLogf(time.Now(), "client stream full, dropping message")
			}
		}
		s.broadcastMtx.RUnlock()
	}
}

func (s *Server) sendBroadcasts(srv pb.RouteChat_StreamServer, token string) {
	stream := s.addSession(token)
	defer s.deleteSesssion(token)

	for {
		select {
		case <-srv.Context().Done():
			return
		case res := <-stream:
			if s, ok := status.FromError(srv.Send(&res)); ok {
				switch s.Code() {
				case codes.OK:
					// noop
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					common.ServerLogf(time.Now(), "client (%s) terminated connection", token)
					return
				default:
					common.ClientLogf(time.Now(), "failed to send to client (%s): %v", token, s.Err())
					return
				}
			}
		}
	}
}
