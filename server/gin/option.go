package gin

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"os"
)

type HandlerFunc func(ctx context.Context, mux *runtime.ServeMux, address string, opts []grpc.DialOption) (err error)
type Option func(s *Server)

func WithHttpHandlerFunc(path string, fs gin.HandlerFunc) Option {
	return func(s *Server) {
		s.HTTPHandlerFuncs[path] = fs
	}
}

func WithHandlerFunc(fs ...HandlerFunc) Option {
	return func(s *Server) {
		s.HandlerFuncs = append(s.HandlerFuncs, fs...)
	}
}

func WithAddress(address string) Option {
	return func(s *Server) {
		s.Address = address
	}
}

func VersionHandler(ctx *gin.Context) {
	version := os.Getenv("VERSION")
	type AppVersion struct {
		Version string `json:"version"`
	}
	if version == "" {
		version = "unknown-version"
	}
	v, _ := json.Marshal(&AppVersion{Version: version})
	_, _ = ctx.Writer.Write(v)
}
