package grpc

import (
	"google.golang.org/grpc"
	LibraryApp "libraryService/internal/libraryApp"
	"net"
)

type Server struct {
	srv *grpc.Server
	lis net.Listener
}

func NewGRPCServer(lis net.Listener, la LibraryApp.App) Server {
	srv := grpc.NewServer()
	RegisterLibraryServiceServer(srv, NewServiceServer(la))
	s := Server{srv: srv, lis: lis}
	return s
}

func (s *Server) SetListener(lis net.Listener) {
	s.lis = lis
}

func (s *Server) Listen() error {
	return s.srv.Serve(s.lis)
}

func (s *Server) Stop() {
	s.srv.Stop()
}

func (s *Server) GetServer() *grpc.Server {
	return s.srv
}
