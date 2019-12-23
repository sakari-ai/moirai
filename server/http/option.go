package http

import (
	"encoding/json"
	"net/http"
	"os"

	"google.golang.org/grpc"
)

type Option func(s *Server)

func WithGRPCAddress(address string) Option {
	return func(s *Server) {
		s.GRPCAddress = address
	}
}

func WithAddress(address string) Option {
	return func(s *Server) {
		s.Address = address
	}
}

func WithDialOption(opts ...grpc.DialOption) Option {
	return func(s *Server) {
		s.DialOptions = append(s.DialOptions, opts...)
	}
}

func WithHandlerFunc(fs ...HandlerFunc) Option {
	return func(s *Server) {
		s.HandlerFuncs = append(s.HandlerFuncs, fs...)
	}
}

func WithHttpHandlerFunc(path string, fs http.HandlerFunc) Option {
	return func(s *Server) {
		s.HTTPHandlerFuncs[path] = fs
	}
}

func VersionHandler(writer http.ResponseWriter, _ *http.Request) {
	version := os.Getenv("VERSION")
	type AppVersion struct {
		Version string `json:"version"`
	}
	if version == "" {
		version = "unknown-version"
	}
	v, _ := json.Marshal(&AppVersion{Version: version})
	_, _ = writer.Write(v)
}
