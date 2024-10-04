package rpc

import (
	"catworks/luna/session/internal/config"
	"catworks/luna/session/pkg/protogo"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type Server struct {
	Srv    *grpc.Server
	config *config.Config
	logger *logrus.Logger

	protogo.UnimplementedSessionServiceServer
}

func NewServer(config *config.Config, logger *logrus.Logger) *Server {
	s := &Server{
		config: config,
		logger: logger,
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
	protogo.RegisterSessionServiceServer(s.Srv, s)
}

func (s Server) Start() error {
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.config.Grpc.Port))
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

func (s Server) StartSession(ctx context.Context, rq *protogo.StartSessionRq) (*protogo.SessionData, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s Server) GetCurrentSession(ctx context.Context, empty *emptypb.Empty) (*protogo.SessionData, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s Server) RenameSession(ctx context.Context, rq *protogo.RenameSessionRq) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s Server) ListSessions(ctx context.Context, empty *emptypb.Empty) (*protogo.SessionList, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s Server) Logout(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s Server) GetInfo(_ context.Context, _ *emptypb.Empty) (*protogo.ServiceInfo, error) {
	return &protogo.ServiceInfo{
		Name:    "luna.session",
		Version: s.config.Version,
	}, nil
}

func (s Server) mustEmbedUnimplementedSessionServiceServer() {
	//TODO implement me
	panic("implement me")
}
