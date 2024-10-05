package rpc

import (
	"catworks/luna/session/pkg/protogo"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s sessionServiceApi) StartSession(ctx context.Context, rq *protogo.StartSessionRq) (*protogo.SessionData, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s sessionServiceApi) GetCurrentSession(ctx context.Context, empty *emptypb.Empty) (*protogo.SessionData, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s sessionServiceApi) RenameSession(ctx context.Context, rq *protogo.RenameSessionRq) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s sessionServiceApi) ListSessions(ctx context.Context, empty *emptypb.Empty) (*protogo.SessionList, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s sessionServiceApi) Logout(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (s sessionServiceApi) GetInfo(_ context.Context, _ *emptypb.Empty) (*protogo.ServiceInfo, error) {
	return &protogo.ServiceInfo{
		Name:    "luna.session",
		Version: s.config.Version,
	}, nil
}
