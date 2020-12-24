package socket

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/Hoongeun/gogoatalk/common"
	"github.com/Hoongeun/gogoatalk/common/util"
	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type SocketListener interface {
	OnUserLogin(userid string, username string, present []*pb.Present)
	OnOtherUserLogin(userid string, username string)
	OnUserLogout(userid string, username string)
	OnSendMessage(message pb.Message)
	OnReadLatest(messages []*pb.Message)
	OnReadMore(messages []*pb.Message)
	OnServerShutdown()
}

type Socket struct {
	conn   *grpc.ClientConn
	client pb.RouteChatClient
	token  string
	sigctx context.Context
	sl     SocketListener
}

func NewSocket(ctx context.Context) *Socket {
	return &Socket{
		sigctx: ctx,
	}
}

func (s *Socket) Connect(ctx context.Context, host string) error {
	connCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.DialContext(connCtx, host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.WithMessage(err, "unable to connect")
	}
	s.conn = conn

	s.client = pb.NewRouteChatClient(conn)
	return nil
}

func (s *Socket) Disconnect() {
	if s.conn != nil {
		s.conn.Close()
	}
}

func (s *Socket) RegisterSocketListener(sl SocketListener) {
	s.sl = sl
}

func (s *Socket) OnEnterChatroom(ctx context.Context) {
	go s.stream(ctx)
}

func (s *Socket) stream(ctx context.Context) error {
	md := metadata.New(map[string]string{"authorization": s.token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	streamClient, err := s.client.Stream(ctx)

	if err != nil {
		return err
	}
	defer streamClient.CloseSend()

	return s.receive(streamClient)
}

func (s *Socket) receive(sc pb.RouteChat_StreamClient) error {
	for {
		res, err := sc.Recv()
		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			common.ClientLogf(time.Now(), "stream canceled (usually indicates shutdown)")
			return nil
		} else if err == io.EOF {
			common.ClientLogf(time.Now(), "stream closed by server")
			return nil
		} else if err != nil {
			return err
		}
		ts := util.TsToTime(res.Timestamp)
		switch evt := res.Event.(type) {
		case *pb.StreamResponse_UserLogin:
			s.sl.OnOtherUserLogin(evt.UserLogin.Userid, evt.UserLogin.Username)
		case *pb.StreamResponse_UserLogout:
			s.sl.OnUserLogout(evt.UserLogout.Userid, evt.UserLogout.Username)
		case *pb.StreamResponse_SendMessage_:
			m := pb.Message{
				Id:        evt.SendMessage.Id,
				Userid:    evt.SendMessage.Userid,
				Text:      evt.SendMessage.Text,
				CreatedAt: evt.SendMessage.CreatedAt,
				UpdatedAt: evt.SendMessage.UpdatedAt,
			}
			s.sl.OnSendMessage(m)
		case *pb.StreamResponse_ServerShutdown:
			s.sl.OnServerShutdown()
			os.Exit(1)
		default:
			common.ClientLogf(ts, "unexpected event from the server: %T", evt)
			return nil
		}
	}
}
