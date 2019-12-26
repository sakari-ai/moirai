package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"
)

type GRPCServer interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context)
}

type GinServer interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context)
}

type HTTPServer interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context)
}

type Server struct {
	GRPCServer GRPCServer `inject:"grpc"`
	HTTPServer HTTPServer `inject:"http"`
	GinServer  GinServer
}

func New(opts ...Option) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Run(ctx context.Context) error {
	s.Start(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-c:
		log.Info(fmt.Sprintf("received signal: %s", s))
	case <-ctx.Done():
		log.Info("received context done")
	}

	s.Stop(ctx)

	return nil
}

func (s *Server) Start(ctx context.Context) {
	var flag uint8
	if s.GRPCServer != nil {
		flag = 1
		go s.startGRPCServer(ctx)
	}
	if s.GinServer != nil {
		flag = 1
		go s.startGinServer(ctx)
	}
	if s.HTTPServer != nil {
		flag = 1
		go s.startHTTPServer(ctx)
	}
	if flag == 0 {
		log.Fatal("what to run huh?", field.Any("server", s))
	}
}

func (s *Server) startGRPCServer(ctx context.Context) {
	if err := s.GRPCServer.Run(ctx); err != nil {
		panic(err)
	}
}

func (s *Server) startHTTPServer(ctx context.Context) {
	if err := s.HTTPServer.Run(ctx); err != nil {
		panic(err)
	}
}

func (s *Server) startGinServer(ctx context.Context) {
	if err := s.GinServer.Run(ctx); err != nil {
		panic(err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	if s.HTTPServer != nil {
		s.HTTPServer.Stop(ctx)
	}
	if s.GRPCServer != nil {
		s.GRPCServer.Stop(ctx)
	}
	if s.GinServer != nil {
		s.GinServer.Stop(ctx)
	}
}
