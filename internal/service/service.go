package service

import (
	"context"

	proto "github.com/ultimathul3/sea-battle-server/proto"
)

type Service struct {
	proto.UnimplementedGameServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateGame(ctx context.Context, req *proto.CreateGameRequest) (*proto.CreateGameResponse, error) {
	return &proto.CreateGameResponse{}, nil
}
