package rpc

import (
	"catworks/luna/session/internal/config"
	"catworks/luna/session/pkg/protogo"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Srv *grpc.Server
	cfg config.GrpcConfig

	session *sessionServiceApi
}

type sessionServiceApi struct {
	config *config.Config
	logger *logrus.Logger

	protogo.UnimplementedSessionServiceServer
}

func NewServer(cfg *config.Config, logger *logrus.Logger) *Server {
	s := &Server{
		cfg: cfg.Grpc,
		session: &sessionServiceApi{
			config: cfg,
			logger: logger,
		},
	}
	s.Srv = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
		),
		grpc.ChainStreamInterceptor(
			grpc_logrus.StreamServerInterceptor(logrus.NewEntry(logger)),
		),
	)

	return s
}

func (s Server) Register() {
	protogo.RegisterSessionServiceServer(s.Srv, s.session)
}

func (s Server) Start() error {
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.cfg.Port))
	if err != nil {
		return err
	}

	err = s.Srv.Serve(l)
	if err != nil {
		return err
	}

	return nil
}

func (s Server) Stop() {
	s.Srv.Stop()
}
