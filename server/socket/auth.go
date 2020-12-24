package socket

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Hoongeun/gogoatalk/protobuf"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type jwtClaims struct {
	Userid string `json:"userid"`
	jwt.StandardClaims
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userid, err := s.am.ValidateUserAccount(req.Username, req.Password)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &jwtClaims{
		Userid: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	tokenString, err := token.SignedString(s.jwtPrivateKey)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	s.am.AppendPresents(userid)
	ps := s.am.GetPresents()

	a, err := s.am.GetAccount(userid)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	s.broadcaster <- pb.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &pb.StreamResponse_UserLogin{&pb.StreamResponse_Login{
			Username: a.Username,
			Userid:   a.Userid,
		}},
	}

	return &pb.LoginResponse{
		Userid:   userid,
		Token:    tokenString,
		Presents: ps,
	}, nil
}

func (s *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.GeneralResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	jwtToken, ok := md["authorization"]
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	c, err := s.validateToken(jwtToken[0])
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}
	p, err := s.am.GetPresent(c.Userid)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "Not presented user.")
	}
	username := p.Username
	userid := p.Userid
	s.am.DeletePresent(c.Userid)

	s.broadcaster <- pb.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &pb.StreamResponse_UserLogout{&pb.StreamResponse_Logout{
			Username: username,
			Userid:   userid,
		}},
	}

	return &pb.GeneralResponse{
		Status:  pb.GeneralResponse_SUCCESS,
		Message: "",
	}, nil
}

func (s *Server) validateToken(token string) (*jwtClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			log.Printf("Unexpected signing method: %v", t.Header["alg"])
			return nil, fmt.Errorf("invalid token")
		}
		return s.jwtPublicKey, nil
	})
	if err != nil && !jwtToken.Valid {
		return nil, err
	}
	if c, ok := jwtToken.Claims.(*jwtClaims); ok {
		return c, nil
	}
	return nil, err
}

func (s *Server) addSession(token string) (sessionChan chan pb.StreamResponse) {
	sessionChan = make(chan pb.StreamResponse, 100)

	s.sessionMtx.Lock()
	s.sessionChans[token] = sessionChan
	s.sessionMtx.Unlock()

	return sessionChan
}

func (s *Server) deleteSesssion(token string) {
	s.sessionMtx.Lock()

	if stream, ok := s.sessionChans[token]; ok {
		delete(s.sessionChans, token)
		close(stream)
	}

	s.sessionMtx.Unlock()
}
