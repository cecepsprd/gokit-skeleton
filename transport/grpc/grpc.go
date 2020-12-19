package grpc

import (
	"log"
	"net"
	"os"
	"os/signal"

	v1 "github.com/cecepsprd/gokit-skeleton/api/proto/v1"
	"github.com/cecepsprd/gokit-skeleton/commons/config"
	"github.com/cecepsprd/gokit-skeleton/internal/endpoint"
	"github.com/cecepsprd/gokit-skeleton/internal/service"
	"google.golang.org/grpc"
)

func RunServer(personSvc service.PersonService, cfg config.Config) error {
	var (
		personEndpoint = endpoint.MakePersonEndpoint(personSvc)
	)

	grpcServer := NewGRPCServer(personEndpoint, nil)

	listen, err := net.Listen("tcp", cfg.App.GrpcPort)
	if err != nil {
		return err
	}

	baseGrpcServer := grpc.NewServer()

	v1.RegisterPersonServer(baseGrpcServer, grpcServer)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			baseGrpcServer.GracefulStop()
		}
	}()

	// start gRPC server
	log.Println("starting Grpc server on port", cfg.App.GrpcPort)
	return baseGrpcServer.Serve(listen)
}
