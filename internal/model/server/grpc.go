package server

import (
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCAPIServer struct {
	*grpc.Server
	address string
}

func NewGRPCServer() *GRPCAPIServer {
	return &GRPCAPIServer{
		address: "127.0.0.1",
	}
}

func (s *GRPCAPIServer) Run() {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		logrus.Fatalf("failed to listen: %s", err.Error())
	}

	go func() {
		if err := s.Serve(listen); err != nil {
			logrus.Fatalf("failed to start grpc server: %s", err.Error())
		}
	}()

	logrus.Infof("start grpc server at %s", s.address)
}

func (s *GRPCAPIServer) Close() {
	s.GracefulStop()
	logrus.Infof("GRPC server on %s stopped", s.address)
}
