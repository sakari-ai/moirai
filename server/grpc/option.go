package grpc

import (
	"google.golang.org/grpc"
)

type Option func(s *Server)

func WithAddress(address string) Option {
	return func(s *Server) {
		s.Address = address
	}
}

func WithServerOption(opts ...grpc.ServerOption) Option {
	return func(s *Server) {
		s.ServerOptions = append(s.ServerOptions, opts...)
	}
}

func WithUnaryInterceptor(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(s *Server) {
		s.UnaryServerInterceptors = append(s.UnaryServerInterceptors, interceptors...)
	}
}

func WithStreamInterceptor(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(s *Server) {
		s.StreamServerInterceptors = append(s.StreamServerInterceptors, interceptors...)
	}
}

type RegisterFunc func(s *grpc.Server)

func WithRegisterFunc(opts ...RegisterFunc) Option {
	return func(s *Server) {
		s.RegisterFuncs = append(s.RegisterFuncs, opts...)
	}
}
