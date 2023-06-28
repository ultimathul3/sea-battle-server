package app

import (
	"fmt"
	"net"

	"github.com/ultimathul3/sea-battle-server/internal/config"
	proto "github.com/ultimathul3/sea-battle-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGrpcServer(cfg *config.Config) error {
	server := grpc.NewServer(grpc.MaxConcurrentStreams(cfg.GRPC.MaxConcurrentStreams))

	proto.RegisterGameServiceServer(server, nil)

	reflection.Register(server)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.IP, cfg.GRPC.Port))
	if err != nil {
		return err
	}

	return server.Serve(listener)
}
