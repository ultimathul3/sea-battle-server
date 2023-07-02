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
	rw     sync.RWMutex
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

func (s *Redis) CreateGame(ctx context.Context, input domain.CreateGameDTO) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if err := s.client.SAdd(ctx, gamesSetKey, input.HostNickname).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, hostFieldKey), input.HostField, keysTTL).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, uuidKey), input.UUID, keysTTL).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(input.HostNickname, statusKey), GameCreated, keysTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Redis) GetGames(ctx context.Context) ([]string, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	var games []string

	members, err := s.client.SMembers(ctx, gamesSetKey).Result()
	if err != nil {
		return nil, err
	}

	for _, member := range members {
		statusValue, err := s.client.Get(ctx, getGameKey(member, statusKey)).Result()
		status, _ := strconv.Atoi(statusValue)

		if err != nil || status != GameCreated {
			s.client.SRem(ctx, gamesSetKey, member)
		} else {
			games = append(games, member)
		}
	}

	return games, nil
}
