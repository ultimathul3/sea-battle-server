package app

import (
	"fmt"
	"net"

	"github.com/ultimathul3/sea-battle-server/internal/config"
	"github.com/ultimathul3/sea-battle-server/internal/service"
	proto "github.com/ultimathul3/sea-battle-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGrpcServer(cfg *config.Config, service *service.Service) error {
	server := grpc.NewServer(grpc.MaxConcurrentStreams(cfg.GRPC.MaxConcurrentStreams))

	proto.RegisterGameServiceServer(server, service)
	reflection.Register(server)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.IP, cfg.GRPC.Port))
	if err != nil {
		return err
	}

	return server.Serve(listener)
}
