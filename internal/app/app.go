package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ultimathul3/sea-battle-server/internal/config"
	"github.com/ultimathul3/sea-battle-server/internal/service"
)

func Run(config *config.Config) error {
	err := make(chan error)

	service := service.New()
	ctx := context.Background()

	go func() {
		err <- runGrpcServer(config)
	}()

	go func() {
		err <- runRestServer(ctx, service, config)
	}()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		err <- nil
	}()

	return <-err
}
