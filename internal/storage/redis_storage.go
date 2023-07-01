package storage

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ultimathul3/sea-battle-server/internal/domain"
)

const (
	keysTTL = 10 * time.Minute
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

func (s *Redis) CreateGame(ctx context.Context, game domain.Game) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if err := s.client.SAdd(ctx, gamesSetKey, game.HostNickname).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(game.HostNickname, fieldKey), game.Field, keysTTL).Err(); err != nil {
		return err
	}
	if err := s.client.Set(ctx, getGameKey(game.HostNickname, uuidKey), game.UUID, keysTTL).Err(); err != nil {
		return err
	}

	return nil
}
