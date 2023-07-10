package storage

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ultimathul3/sea-battle-server/internal/domain"
)

const (
	keysTTL = 20 * time.Minute
)

type Redis struct {
	client *redis.Client
	mutex  sync.Mutex
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

func (s *Redis) CreateGame(ctx context.Context, input domain.CreateGameDTO) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if err := s.client.SAdd(ctx, gamesSetKey, input.HostNickname).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, hostFieldKey), input.HostField, keysTTL).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, hostUuidKey), input.HostUuid, keysTTL).Err(); err != nil {
		return err
	}
	if err := s.SetStatus(ctx, domain.SetStatusDTO{
		HostNickname: input.HostNickname,
		Status:       domain.GameCreatedStatus,
		KeepTTL:      false,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Redis) GetGames(ctx context.Context) ([]string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var games []string

	members, err := s.client.SMembers(ctx, gamesSetKey).Result()
	if err != nil {
		return nil, err
	}

	for _, member := range members {
		status, err := s.GetStatus(ctx, domain.GetStatusDTO{HostNickname: member})

		if err != nil || status != domain.GameCreatedStatus {
			s.client.SRem(ctx, gamesSetKey, member)
		} else {
			games = append(games, member)
		}
	}

	return games, nil
}

func (s *Redis) JoinGame(ctx context.Context, input domain.JoinGameDTO) error {
	status, err := s.GetStatus(ctx, domain.GetStatusDTO{
		HostNickname: input.HostNickname,
	})
	if err != nil {
		return err
	}

	if status != domain.GameCreatedStatus {
		return domain.ErrGameDoesNotExist
	}

	if err := s.SetStatus(ctx, domain.SetStatusDTO{
		HostNickname: input.HostNickname,
		Status:       domain.WaitingForOpponentStatus,
		KeepTTL:      true,
	}); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, opponentNicknameKey), input.OpponentNickname, keysTTL).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, opponentUuidKey), input.OpponentUuid, keysTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Redis) StartGame(ctx context.Context, input domain.StartGameDTO) error {
	status, err := s.GetStatus(ctx, domain.GetStatusDTO{
		HostNickname: input.HostNickname,
	})
	if err != nil {
		return err
	}

	if status != domain.WaitingForOpponentStatus {
		return domain.ErrGameDoesNotExist
	}

	opponentUuid, err := s.client.Get(ctx, getGameKey(input.HostNickname, opponentUuidKey)).Result()
	if err != nil {
		return domain.ErrGameDoesNotExist
	}

	if opponentUuid != input.OpponentUuid {
		return domain.ErrGameDoesNotExist
	}

	if err := s.client.Set(ctx, getGameKey(input.HostNickname, opponentFieldKey), input.OpponentField, keysTTL).Err(); err != nil {
		return err
	}
	if err := s.SetStatus(ctx, domain.SetStatusDTO{
		HostNickname: input.HostNickname,
		Status:       domain.GameStartedStatus,
		KeepTTL:      true,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Redis) GetRoleByUuid(ctx context.Context, input domain.GetRoleByUuidDTO) (domain.Role, error) {
	opponentUuid, err := s.client.Get(ctx, getGameKey(input.HostNickname, opponentUuidKey)).Result()
	if err != nil {
		return domain.UnknownRole, domain.ErrGameDoesNotExist
	}

	if opponentUuid != input.Uuid {
		hostUuid, err := s.client.Get(ctx, getGameKey(input.HostNickname, hostUuidKey)).Result()
		if err != nil {
			return domain.UnknownRole, domain.ErrGameDoesNotExist
		}

		if hostUuid != input.Uuid {
			return domain.UnknownRole, domain.ErrGameDoesNotExist
		}

		return domain.HostRole, nil
	}

	return domain.OpponentRole, nil
}

func (s *Redis) GetStatus(ctx context.Context, input domain.GetStatusDTO) (domain.Status, error) {
	statusValue, err := s.client.Get(ctx, getGameKey(input.HostNickname, statusKey)).Result()
	if err != nil {
		return domain.UnknownStatus, domain.ErrGameDoesNotExist
	}

	status, err := strconv.Atoi(statusValue)
	if err != nil {
		return domain.UnknownStatus, err
	}

	return domain.Status(status), nil
}

func (s *Redis) SetStatus(ctx context.Context, input domain.SetStatusDTO) error {
	if input.KeepTTL {
		return s.client.Set(ctx, getGameKey(input.HostNickname, statusKey), uint8(input.Status), redis.KeepTTL).Err()
	}
	return s.client.Set(ctx, getGameKey(input.HostNickname, statusKey), uint8(input.Status), keysTTL).Err()
}

func (s *Redis) GetField(ctx context.Context, input domain.GetFieldDTO) (string, error) {
	switch input.Role {
	case domain.HostRole:
		field, err := s.client.Get(ctx, getGameKey(input.HostNickname, hostFieldKey)).Result()
		if err != nil {
			return "", err
		}
		return field, nil
	case domain.OpponentRole:
		field, err := s.client.Get(ctx, getGameKey(input.HostNickname, opponentFieldKey)).Result()
		if err != nil {
			return "", err
		}
		return field, nil
	default:
		return "", domain.ErrGameDoesNotExist
	}
}

func (s *Redis) UpdateField(ctx context.Context, input domain.UpdateFieldDTO) error {
	switch input.Role {
	case domain.HostRole:
		if err := s.client.Set(ctx, getGameKey(input.HostNickname, hostFieldKey), input.Field, redis.KeepTTL).Err(); err != nil {
			return err
		}
		return nil
	case domain.OpponentRole:
		if err := s.client.Set(ctx, getGameKey(input.HostNickname, opponentFieldKey), input.Field, redis.KeepTTL).Err(); err != nil {
			return err
		}
		return nil
	default:
		return domain.ErrGameDoesNotExist
	}
}
