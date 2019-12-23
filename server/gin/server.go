package gin

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	Mux              *runtime.ServeMux
	GRPCAddress      string
	DialOptions      []grpc.DialOption
	Address          string
	HTTPHandlerFuncs map[string]gin.HandlerFunc
	HandlerFuncs     []HandlerFunc
	ginHandler       *gin.Engine
	HttpServer       *http.Server
}

func New(opts ...Option) *Server {
	s := &Server{
		DialOptions:      []grpc.DialOption{grpc.WithInsecure()},
		ginHandler:       gin.Default(),
		HTTPHandlerFuncs: make(map[string]gin.HandlerFunc),
		Mux:              runtime.NewServeMux(),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Run(ctx context.Context) error {
	for _, f := range s.HandlerFuncs {
		if err := f(ctx, s.Mux, s.GRPCAddress, s.DialOptions); err != nil {
			return err
		}
	}
	for path, hF := range s.HTTPHandlerFuncs {
		s.ginHandler.Any(path, hF)
	}
	s.HttpServer = &http.Server{
		Addr:    s.Address,
		Handler: s.ginHandler,
	}
	return s.Start(ctx)
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.HttpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.HttpServer.Shutdown(ctx); err != nil {
		log.Fatal("ginHandler Shutdown: ", err)
	}
}
