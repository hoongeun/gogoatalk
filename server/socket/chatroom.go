package socket

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func (s *Server) ReadLatest(ctx context.Context, req *pb.ReadLatestRequest) (*pb.ReadGeneralResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	jwtToken, ok := md["authorization"]
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	_, err := s.validateToken(jwtToken[0])
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	messages, err := s.cm.ReadLatest()
	if err != nil {
		return nil, grpc.Errorf(codes.Unavailable, "Fail to ReadLatest()")
	}

	pbms := make([]*pb.Message, 0)

	for _, m := range messages {
		pbm := &pb.Message{
			Id:        m.Id,
			Text:      m.Text,
			Userid:    m.Userid,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
		pbms = append(pbms, pbm)
	}

	return &pb.ReadGeneralResponse{
		Messages: pbms,
	}, nil
}

func (s *Server) ReadMore(ctx context.Context, req *pb.ReadMoreRequest) (*pb.ReadGeneralResponse, error) {
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

	_, err = s.am.GetPresent(c.Userid) // no more error
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "Not presented user.")
	}

	messages, err := s.cm.ReadMore(req.Id, int(req.More))
	if err != nil {
		return nil, grpc.Errorf(codes.Unavailable, "Fail to ReadMore()")
	}

	pbms := make([]*pb.Message, 0)

	for _, m := range messages {
		pbm := &pb.Message{
			Id:        m.Id,
			Text:      m.Text,
			Userid:    m.Userid,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
		pbms = append(pbms, pbm)
	}

	return &pb.ReadGeneralResponse{
		Messages: pbms,
	}, nil
}

func (s *Server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.GeneralResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	jwtToken, ok := md["authorization"]
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	if len(jwtToken) == 0 || jwtToken[0] == "" {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}
	c, err := s.validateToken(jwtToken[0])
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}
	p, err := s.am.GetPresent(c.Userid)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}
	m, err := s.cm.Append(p.Userid, req.Text)
	if err != nil {
		return nil, grpc.Errorf(codes.Unavailable, "Fail to SendMessage()")
	}

	s.broadcaster <- pb.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &pb.StreamResponse_SendMessage_{&pb.StreamResponse_SendMessage{
			Id:        m.Id,
			Text:      m.Text,
			Userid:    m.Userid,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}},
	}

	return &pb.GeneralResponse{
		Status: pb.GeneralResponse_SUCCESS,
	}, nil
}

func (s *Server) RemoveMessage(ctx context.Context, req *pb.RemoveMessageRequest) (*pb.GeneralResponse, error) {
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
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	m, err := s.cm.Remove(p.Userid, req.Id)
	if err != nil {
		return nil, grpc.Errorf(codes.Unavailable, "Fail to SendMessage()")
	}

	s.broadcaster <- pb.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &pb.StreamResponse_RemoveMessage_{&pb.StreamResponse_RemoveMessage{
			Id: m.Id,
		}},
	}

	return &pb.GeneralResponse{
		Status: pb.GeneralResponse_SUCCESS,
	}, nil
}

func (s *Server) UpdateMessage(ctx context.Context, req *pb.UpdateMessageRequest) (*pb.GeneralResponse, error) {
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

	a, err := s.am.GetPresent(c.Userid)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	m, err := s.cm.Update(a.Userid, req.Id, req.Text)
	if err != nil {
		return nil, grpc.Errorf(codes.Unavailable, "Fail to SendMessage()")
	}

	s.broadcaster <- pb.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &pb.StreamResponse_UpdateMessage_{&pb.StreamResponse_UpdateMessage{
			Id:        m.Id,
			Text:      m.Text,
			UpdatedAt: m.UpdatedAt,
		}},
	}

	return &pb.GeneralResponse{
		Status: pb.GeneralResponse_SUCCESS,
	}, nil
}
