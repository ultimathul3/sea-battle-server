package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ultimathul3/sea-battle-server/internal/config"
	"github.com/ultimathul3/sea-battle-server/internal/service"
	proto "github.com/ultimathul3/sea-battle-server/proto"
)

func runRestServer(ctx context.Context, service *service.Service, cfg *config.Config) error {
	gmux := runtime.NewServeMux()

	err := proto.RegisterGameServiceHandlerServer(ctx, gmux, service)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gmux)
	fs := http.FileServer(http.Dir("./swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	return http.ListenAndServe(
		fmt.Sprintf("%s:%d", cfg.HTTP.IP, cfg.HTTP.Port),
		mux,
	)
}
