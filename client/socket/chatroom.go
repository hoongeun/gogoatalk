package socket

import (
	"context"
	"time"

	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"google.golang.org/grpc/metadata"
)

func (s *Socket) ReadLatest(ctx context.Context) ([]*pb.Message, error) {
	md := metadata.New(map[string]string{"authorization": s.token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	connCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := s.client.ReadLatest(connCtx, &pb.ReadLatestRequest{})

	if err != nil {
		return nil, err
	}

	s.sl.OnReadLatest(res.Messages)

	return res.Messages, nil
}

func (s *Socket) SendMessage(ctx context.Context, text string) error {
	md := metadata.New(map[string]string{"authorization": s.token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	connCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err := s.client.SendMessage(connCtx, &pb.SendMessageRequest{
		Text: text,
	})
	return err
}
