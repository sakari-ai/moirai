package server

type Option func(s *Server)

func WithGRPCServer(srv GRPCServer) Option {
	return func(s *Server) {
		s.GRPCServer = srv
	}
}

func WithHTTPServer(srv HTTPServer) Option {
	return func(s *Server) {
		s.HTTPServer = srv
	}
}
