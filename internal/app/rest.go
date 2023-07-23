package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ultimathul3/sea-battle-server/internal/config"
	"github.com/ultimathul3/sea-battle-server/internal/service"
	proto "github.com/ultimathul3/sea-battle-server/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func enableCors(origin, methods string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", methods)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func runRestServer(ctx context.Context, service *service.Service, cfg *config.Config) error {
	gmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
		},
	}))

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
		enableCors(cfg.Cors.AllowOrigin, cfg.Cors.AllowMethods, mux),
	)
}
