package http

import (
	"context"
	"fmt"
	"github.com/sakari-ai/moirai/log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type HandlerFunc func(ctx context.Context, mux *runtime.ServeMux, address string, opts []grpc.DialOption) (err error)
type Server struct {
	Mux              *runtime.ServeMux
	httpMux          *http.ServeMux
	Server           *http.Server
	GRPCAddress      string
	Address          string
	DialOptions      []grpc.DialOption
	HandlerFuncs     []HandlerFunc
	HTTPHandlerFuncs map[string]http.HandlerFunc
}

func WhitelistHeaders() runtime.HeaderMatcherFunc {
	return func(key string) (string, bool) {
		k := strings.ToLower(key)
		whitelist := map[string]struct{}{
			"x-accountid":  struct{}{},
			"x-authsakari": struct{}{},
		}
		if _, ok := whitelist[k]; ok {
			return key, true
		}
		return runtime.DefaultHeaderMatcher(key)

	}
}

func New(opts ...Option) *Server {
	s := &Server{
		DialOptions:      []grpc.DialOption{grpc.WithInsecure()},
		HTTPHandlerFuncs: make(map[string]http.HandlerFunc),
		Mux:              runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(WhitelistHeaders())),
	}
	for _, opt := range opts {
		opt(s)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", s.Mux)
	s.httpMux = httpMux
	return s
}

func (s *Server) Run(ctx context.Context) error {
	for _, f := range s.HandlerFuncs {
		if err := f(ctx, s.Mux, s.GRPCAddress, s.DialOptions); err != nil {
			return err
		}
	}
	for path, hF := range s.HTTPHandlerFuncs {
		s.httpMux.HandleFunc(path, hF)
	}
	s.Server = &http.Server{
		Addr:    s.Address,
		Handler: s.httpMux,
	}
	log.Info(fmt.Sprintf(`HTTP listening on: "%s" (GRPC at: "%s")`, s.Address, s.GRPCAddress))
	return s.Start(ctx)
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) {
	log.Info("HTTP shutting down...")
	s.Server.Shutdown(ctx)
	log.Info("HTTP gracefully stopped")
}
