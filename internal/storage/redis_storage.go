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
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, statusKey), uint8(domain.GameCreated), keysTTL).Err(); err != nil {
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
		statusValue, err := s.client.Get(ctx, getGameKey(member, statusKey)).Result()
		status, _ := strconv.Atoi(statusValue)

		if err != nil || domain.Status(status) != domain.GameCreated {
			s.client.SRem(ctx, gamesSetKey, member)
		} else {
			games = append(games, member)
		}
	}

	return games, nil
}

func (s *Redis) JoinGame(ctx context.Context, input domain.JoinGameDTO) error {
	statusValue, err := s.client.Get(ctx, getGameKey(input.HostNickname, statusKey)).Result()
	status, _ := strconv.Atoi(statusValue)
	if err != nil {
		return domain.ErrGameDoesNotExist
	}

	if domain.Status(status) != domain.GameCreated {
		return domain.ErrGameDoesNotExist
	}

	if err := s.client.Set(ctx, getGameKey(input.HostNickname, statusKey), uint8(domain.WaitingForOpponent), redis.KeepTTL).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, opponentNicknameKey), input.OpponentNickname, redis.KeepTTL).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, opponentUuidKey), input.OpponentUuid, redis.KeepTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Redis) StartGame(ctx context.Context, input domain.StartGameDTO) error {
	opponentUuid, err := s.client.Get(ctx, getGameKey(input.HostNickname, opponentUuidKey)).Result()
	if err != nil {
		return domain.ErrGameDoesNotExist
	}

	if opponentUuid != input.OpponentUuid {
		return domain.ErrGameDoesNotExist
	}

	if err := s.client.Set(ctx, getGameKey(input.HostNickname, opponentFieldKey), input.OpponentField, redis.KeepTTL).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, statusKey), uint8(domain.GameStarted), keysTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Redis) GetFieldForShoot(ctx context.Context, input domain.GetFieldForShootDTO) (string, error) {
	opponentUuid, err := s.client.Get(ctx, getGameKey(input.HostNickname, opponentUuidKey)).Result()
	if err != nil {
		return "", domain.ErrGameDoesNotExist
	}

	if opponentUuid != input.Uuid {
		hostUuid, err := s.client.Get(ctx, getGameKey(input.HostNickname, hostUuidKey)).Result()
		if err != nil {
			return "", domain.ErrGameDoesNotExist
		}

		if hostUuid != input.Uuid {
			return "", domain.ErrGameDoesNotExist
		}
	}

	statusValue, err := s.client.Get(ctx, getGameKey(input.HostNickname, statusKey)).Result()
	status, _ := strconv.Atoi(statusValue)
	if err != nil {
		return "", domain.ErrGameDoesNotExist
	}

	switch domain.Status(status) {
	case domain.GameStarted, domain.OpponentHit, domain.OpponentMiss:
		field, err := s.client.Get(ctx, getGameKey(input.HostNickname, opponentFieldKey)).Result()
		if err != nil {
			return "", err
		}
		return field, nil
	case domain.HostHit, domain.HostMiss:
		field, err := s.client.Get(ctx, getGameKey(input.HostNickname, hostFieldKey)).Result()
		if err != nil {
			return "", err
		}
		return field, nil
	default:
		return "", domain.ErrGameDoesNotExist
	}
}
