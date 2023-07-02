package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ultimathul3/sea-battle-server/internal/domain"
	proto "github.com/ultimathul3/sea-battle-server/proto"
)

type Storage interface {
	CreateGame(ctx context.Context, input domain.CreateGameDTO) error
	GetGames(ctx context.Context) ([]string, error)
}

type Service struct {
	proto.UnimplementedGameServiceServer
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateGame(ctx context.Context, req *proto.CreateGameRequest) (*proto.CreateGameResponse, error) {
	gameUUID := uuid.NewString()

	if err := s.storage.CreateGame(ctx, domain.CreateGameDTO{
		HostNickname: req.HostNickname,
		HostField:    req.HostField,
		UUID:         gameUUID,
	}); err != nil {
		return nil, err
	}

	return &proto.CreateGameResponse{
		UUID: gameUUID,
	}, nil
}

func (s *Service) GetGames(ctx context.Context, req *proto.GetGamesRequest) (*proto.GetGamesResponse, error) {
	games, err := s.storage.GetGames(ctx)

	return &proto.GetGamesResponse{
		Games: games,
	}, err
}
