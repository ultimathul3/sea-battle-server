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
	JoinGame(ctx context.Context, input domain.JoinGameDTO) error
	StartGame(ctx context.Context, input domain.StartGameDTO) error
	GetFieldForShoot(ctx context.Context, input domain.GetFieldForShootDTO) (string, error)
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
	hostUUID := uuid.NewString()

	createGameDTO := domain.CreateGameDTO{
		HostNickname: req.HostNickname,
		HostField:    req.HostField,
		HostUuid:     hostUUID,
	}
	if err := createGameDTO.Validate(); err != nil {
		return nil, err
	}

	if err := s.storage.CreateGame(ctx, createGameDTO); err != nil {
		return nil, err
	}

	return &proto.CreateGameResponse{
		HostUuid: hostUUID,
	}, nil
}

func (s *Service) GetGames(ctx context.Context, req *proto.GetGamesRequest) (*proto.GetGamesResponse, error) {
	games, err := s.storage.GetGames(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.GetGamesResponse{
		Games: games,
	}, nil
}

func (s *Service) JoinGame(ctx context.Context, req *proto.JoinGameRequest) (*proto.JoinGameResponse, error) {
	opponentUuid := uuid.NewString()

	joinGameDTO := domain.JoinGameDTO{
		HostNickname:     req.HostNickname,
		OpponentNickname: req.OpponentNickname,
		OpponentUuid:     opponentUuid,
	}
	if err := joinGameDTO.Validate(); err != nil {
		return nil, err
	}

	if err := s.storage.JoinGame(ctx, joinGameDTO); err != nil {
		return nil, err
	}

	return &proto.JoinGameResponse{
		OpponentUuid: opponentUuid,
	}, nil
}

func (s *Service) StartGame(ctx context.Context, req *proto.StartGameRequest) (*proto.StartGameResponse, error) {
	startGameDTO := domain.StartGameDTO{
		HostNickname:  req.HostNickname,
		OpponentField: req.OpponentField,
		OpponentUuid:  req.OpponentUuid,
	}
	if err := startGameDTO.Validate(); err != nil {
		return nil, err
	}

	if err := s.storage.StartGame(ctx, startGameDTO); err != nil {
		return nil, err
	}

	return &proto.StartGameResponse{}, nil
}

func (s *Service) Shoot(ctx context.Context, req *proto.ShootRequest) (*proto.ShootResponse, error) {
	getFieldForShootDTO := domain.GetFieldForShootDTO{
		HostNickname: req.HostNickname,
		Uuid:         req.Uuid,
	}
	if err := getFieldForShootDTO.Validate(); err != nil {
		return nil, err
	}

	// TODO
	_, err := s.storage.GetFieldForShoot(ctx, getFieldForShootDTO)
	if err != nil {
		return nil, err
	}

	fieldIndex := req.X*domain.FieldDimension + req.Y
	shootDTO := domain.ShootDTO{
		HostNickname: req.HostNickname,
		FieldIndex:   fieldIndex,
	}
	if err := shootDTO.Validate(); err != nil {
		return nil, err
	}

	return &proto.ShootResponse{
		Status: convertStatusToProto(domain.Unknown),
	}, nil
}

func convertStatusToProto(status domain.Status) proto.Status {
	switch status {
	case domain.GameCreated:
		return proto.Status_GAME_CREATED
	case domain.WaitingForOpponent:
		return proto.Status_WAITING_FOR_OPPONENT
	case domain.GameStarted:
		return proto.Status_GAME_STARTED
	case domain.HostHit:
		return proto.Status_HOST_HIT
	case domain.HostMiss:
		return proto.Status_HOST_MISS
	case domain.OpponentHit:
		return proto.Status_OPPONENT_HIT
	case domain.OpponentMiss:
		return proto.Status_OPPONENT_MISS
	case domain.HostWon:
		return proto.Status_HOST_WON
	case domain.OpponentWon:
		return proto.Status_OPPONENT_WON
	default:
		return proto.Status_UNKNOWN
	}
}
