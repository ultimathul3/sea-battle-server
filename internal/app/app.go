package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ultimathul3/sea-battle-server/internal/config"
	"github.com/ultimathul3/sea-battle-server/internal/service"
	"github.com/ultimathul3/sea-battle-server/internal/storage"
	"github.com/ultimathul3/sea-battle-server/pkg/redis"
)

func Run(cfg *config.Config) error {
	errc := make(chan error)
	ctx := context.Background()

	redisClient, err := redis.NewClient(ctx, cfg)
	if err != nil {
		return err
	}

	storage := storage.NewRedis(redisClient)
	service := service.New(storage)

	go func() {
		errc <- runGrpcServer(cfg)
	}()

	go func() {
		errc <- runRestServer(ctx, service, cfg)
	}()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		errc <- nil
	}()

	return <-errc
}
