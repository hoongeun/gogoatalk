package socket

import (
	"context"
	"time"

	"github.com/Hoongeun/gogoatalk/common"
	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Socket) Login(ctx context.Context, username string, password string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := s.client.Login(ctx, &pb.LoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return err
	}

	s.token = res.Token
	s.sl.OnUserLogin(res.Userid, username, res.Presents)

	return nil
}

func (s *Socket) Logout(ctx context.Context) error {
	md := metadata.New(map[string]string{"authorization": s.token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	connCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := s.client.Logout(connCtx, &pb.LogoutRequest{})
	if s, ok := status.FromError(err); ok && s.Code() == codes.Unavailable {
		common.ClientLogf(time.Now(), "unable to logout (connection already closed)")
		return nil
	}

	s.token = ""

	return err
}
