package grpc

import (
	"context"
	"github.com/sakari-ai/moirai/log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

type Server struct {
	ServerOptions            []grpc.ServerOption
	StreamServerInterceptors []grpc.StreamServerInterceptor
	UnaryServerInterceptors  []grpc.UnaryServerInterceptor
	RegisterFuncs            []RegisterFunc
	Server                   *grpc.Server
	Address                  string
}

func New(opts ...Option) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	s.ServerOptions = append(s.ServerOptions,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(s.StreamServerInterceptors...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(s.UnaryServerInterceptors...)),
	)

	s.Server = grpc.NewServer(s.ServerOptions...)

	for _, f := range s.RegisterFuncs {
		f(s.Server)
	}
	return s
}

func (s *Server) Run(ctx context.Context) error {
	listen, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	log.Info("GRPC listening on: " + s.Address)
	if err := s.Server.Serve(listen); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) {
	log.Info("GRPC shutting down...")
	s.Server.GracefulStop()
	log.Info("GRPC gracefully stopped")
}
