package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ultimathul3/sea-battle-server/internal/domain"
	proto "github.com/ultimathul3/sea-battle-server/proto"
)

type Storage interface {
	CreateGame(ctx context.Context, input domain.CreateGameDTO) error
	GetGames(ctx context.Context) ([]string, error)
	JoinGame(ctx context.Context, input domain.JoinGameDTO) error
	StartGame(ctx context.Context, input domain.StartGameDTO) error
	GetRoleByUuid(ctx context.Context, input domain.GetRoleByUuidDTO) (domain.Role, error)
	GetStatus(ctx context.Context, input domain.GetStatusDTO) (domain.Status, error)
	SetStatus(ctx context.Context, input domain.SetStatusDTO) error
	GetField(ctx context.Context, input domain.GetFieldDTO) (string, error)
	UpdateField(ctx context.Context, input domain.UpdateFieldDTO) error
	GetUuid(ctx context.Context, input domain.GetUuidDTO) (string, error)
}

type Service struct {
	proto.UnimplementedGameServiceServer
	storage    Storage
	events     map[string]chan domain.Event
	timestamps map[string]time.Time
}

func New(storage Storage) *Service {
	return &Service{
		storage:    storage,
		events:     make(map[string]chan domain.Event),
		timestamps: make(map[string]time.Time),
	}
}

func (s *Service) ObserveEventsTTL(eventsTTL time.Duration) {
	for range time.Tick(eventsTTL) {
		for uuid, timestamp := range s.timestamps {
			if time.Since(timestamp) > eventsTTL {
				delete(s.events, uuid)
				delete(s.timestamps, uuid)
			}
		}
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

	s.events[hostUUID] = make(chan domain.Event, domain.AllShipsCellsNumber)
	s.timestamps[hostUUID] = time.Now()

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

	s.events[opponentUuid] = make(chan domain.Event, domain.AllShipsCellsNumber)
	s.timestamps[opponentUuid] = time.Now()

	hostUuid, err := s.storage.GetUuid(ctx, domain.GetUuidDTO{
		HostNickname: req.HostNickname,
		Role:         domain.HostRole,
	})
	if err != nil {
		return nil, err
	}

	status, err := s.storage.GetStatus(ctx, domain.GetStatusDTO{
		HostNickname: req.HostNickname,
	})
	if err != nil {
		return nil, err
	}

	s.events[hostUuid] <- domain.Event{
		Status:  status,
		Message: req.OpponentNickname,
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

	hostUuid, err := s.storage.GetUuid(ctx, domain.GetUuidDTO{
		HostNickname: req.HostNickname,
		Role:         domain.HostRole,
	})
	if err != nil {
		return nil, err
	}

	status, err := s.storage.GetStatus(ctx, domain.GetStatusDTO{
		HostNickname: req.HostNickname,
	})
	if err != nil {
		return nil, err
	}

	s.events[hostUuid] <- domain.Event{
		Status: status,
	}

	return &proto.StartGameResponse{}, nil
}

func (s *Service) Shoot(ctx context.Context, req *proto.ShootRequest) (*proto.ShootResponse, error) {
	status, err := s.storage.GetStatus(ctx, domain.GetStatusDTO{
		HostNickname: req.HostNickname,
	})
	if err != nil {
		return nil, err
	}

	getRoleByUuidDTO := domain.GetRoleByUuidDTO{
		HostNickname: req.HostNickname,
		Uuid:         req.Uuid,
	}
	if err := getRoleByUuidDTO.Validate(); err != nil {
		return nil, err
	}

	role, err := s.storage.GetRoleByUuid(ctx, getRoleByUuidDTO)
	if err != nil {
		return nil, err
	}

	switch role {
	case domain.HostRole:
		if status != domain.GameStartedStatus && status != domain.OpponentMissStatus && status != domain.HostHitStatus {
			return nil, domain.ErrInvalidAction
		}
	case domain.OpponentRole:
		if status != domain.HostMissStatus && status != domain.OpponentHitStatus {
			return nil, domain.ErrInvalidAction
		}
	}

	oppositeRole := domain.HostRole
	if role == domain.HostRole {
		oppositeRole = domain.OpponentRole
	}

	getFieldDTO := domain.GetFieldDTO{
		HostNickname: req.HostNickname,
		Role:         oppositeRole,
	}

	field, err := s.storage.GetField(ctx, getFieldDTO)
	if err != nil {
		return nil, err
	}

	var x, y = req.X, req.Y
	shootDTO := domain.ShootDTO{X: x, Y: y}
	if err := shootDTO.Validate(); err != nil {
		return nil, err
	}

	destroyedShip := domain.UnknownShip
	var destroyedX, destroyedY int

	matrixField := domain.ConvertFieldToRuneMatrix(field)
	switch matrixField[y+1][x+1] {
	case domain.EmptyCell, domain.OccupiedCell:
		matrixField[y+1][x+1] = domain.MissCell
		if role == domain.HostRole {
			status = domain.HostMissStatus
		} else {
			status = domain.OpponentMissStatus
		}
	case domain.SingleDeckShipCell,
		domain.DoubleDeckShipDownCell, domain.ThreeDeckShipDownCell, domain.FourDeckShipDownCell,
		domain.DoubleDeckShipRightCell, domain.ThreeDeckShipRightCell, domain.FourDeckShipRightCell,
		domain.ShipLeftCell, domain.ShipUpCell, domain.ShipLeftEndCell, domain.ShipUpEndCell:
		destroyedShip, destroyedX, destroyedY = domain.Shoot(matrixField, int(y), int(x))

		isWin := domain.IsWin(matrixField)

		if role == domain.HostRole {
			status = domain.HostHitStatus
			if isWin {
				status = domain.HostWonStatus
			}
		} else {
			status = domain.OpponentHitStatus
			if isWin {
				status = domain.OpponentWonStatus
			}
		}
	default:
		return nil, domain.ErrInvalidHitCell
	}

	updateFieldDTO := domain.UpdateFieldDTO{
		HostNickname: req.HostNickname,
		Role:         oppositeRole,
		Field:        domain.ConvertFieldRuneMatrixToString(matrixField),
	}
	if err := s.storage.UpdateField(ctx, updateFieldDTO); err != nil {
		return nil, err
	}

	if err := s.storage.SetStatus(ctx, domain.SetStatusDTO{
		HostNickname: req.HostNickname,
		Status:       status,
	}); err != nil {
		return nil, err
	}

	oppositeRoleUuid, err := s.storage.GetUuid(ctx, domain.GetUuidDTO{
		HostNickname: req.HostNickname,
		Role:         oppositeRole,
	})
	if err != nil {
		return nil, err
	}

	s.events[oppositeRoleUuid] <- domain.Event{
		Status: status,
		X:      x,
		Y:      y,
	}

	if destroyedShip != domain.UnknownShip {
		return &proto.ShootResponse{
			Status:        convertStatusToProto(status),
			X:             uint32(destroyedX),
			Y:             uint32(destroyedY),
			DestroyedShip: convertShipToProto(destroyedShip),
		}, nil
	}

	return &proto.ShootResponse{
		Status: convertStatusToProto(status),
	}, nil
}

func (s *Service) Wait(ctx context.Context, req *proto.WaitRequest) (*proto.WaitResponse, error) {
	ch, ok := s.events[req.Uuid]
	if !ok {
		return nil, domain.ErrGameDoesNotExist
	}

	event := <-ch

	return &proto.WaitResponse{
		Status:  convertStatusToProto(event.Status),
		X:       event.X,
		Y:       event.Y,
		Message: event.Message,
	}, nil
}

func convertStatusToProto(status domain.Status) proto.Status {
	switch status {
	case domain.GameCreatedStatus:
		return proto.Status_GAME_CREATED
	case domain.WaitingForOpponentStatus:
		return proto.Status_WAITING_FOR_OPPONENT
	case domain.GameStartedStatus:
		return proto.Status_GAME_STARTED
	case domain.HostHitStatus:
		return proto.Status_HOST_HIT
	case domain.HostMissStatus:
		return proto.Status_HOST_MISS
	case domain.OpponentHitStatus:
		return proto.Status_OPPONENT_HIT
	case domain.OpponentMissStatus:
		return proto.Status_OPPONENT_MISS
	case domain.HostWonStatus:
		return proto.Status_HOST_WON
	case domain.OpponentWonStatus:
		return proto.Status_OPPONENT_WON
	default:
		return proto.Status_UNKNOWN_STATUS
	}
}

func convertShipToProto(ship domain.Ship) proto.Ship {
	switch ship {
	case domain.SingleDeckShip:
		return proto.Ship_SINGLE_DECK
	case domain.DoubleDeckShipDown:
		return proto.Ship_DOUBLE_DECK_DOWN
	case domain.ThreeDeckShipDown:
		return proto.Ship_THREE_DECK_DOWN
	case domain.FourDeckShipDown:
		return proto.Ship_FOUR_DECK_DOWN
	case domain.DoubleDeckShipRight:
		return proto.Ship_DOUBLE_DECK_RIGHT
	case domain.ThreeDeckShipRight:
		return proto.Ship_THREE_DECK_RIGHT
	case domain.FourDeckShipRight:
		return proto.Ship_FOUR_DECK_RIGHT
	default:
		return proto.Ship_UNKNOWN_SHIP
	}
}
