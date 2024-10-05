package rpc

import (
	"catworks/luna/session/internal/domain"
	"catworks/luna/session/pkg/protogo"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *sessionServiceApi) StartSession(ctx context.Context, rq *protogo.StartSessionRq) (*protogo.SessionData, error) {
	var deviceType domain.DeviceType
	switch rq.DeviceType {
	case protogo.DeviceType_MOBILE:
		deviceType = domain.MOBILE
	case protogo.DeviceType_TV:
		deviceType = domain.TV
	}

	usecaseRq := domain.CreateSessionRq{
		Type: deviceType,
		Name: rq.Name,
	}

	session, err := s.sessionUc.Create(ctx, &usecaseRq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return s.sessionToRpc(session), nil
}

func (s *sessionServiceApi) GetCurrentSession(ctx context.Context, _ *emptypb.Empty) (*protogo.SessionData, error) {
	session, err := s.sessionFromAuth(ctx)
	if err != nil {
		return nil, err
	}

	return s.sessionToRpc(session), nil
}

func (s *sessionServiceApi) RenameSession(ctx context.Context, rq *protogo.RenameSessionRq) (*emptypb.Empty, error) {
	session, err := s.sessionFromAuth(ctx)
	if err != nil {
		return nil, err
	}

	usecaseRq := &domain.RenameSessionRq{
		Name: rq.Name,
		Id:   session.Id,
	}

	ref := rq.Session
	if ref != nil {
		usecaseRq.Id = ref.SessionId
	}

	if err := s.sessionUc.Rename(ctx, usecaseRq); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *sessionServiceApi) ListSessions(ctx context.Context, _ *emptypb.Empty) (*protogo.SessionList, error) {
	if _, err := s.sessionFromAuth(ctx); err != nil {
		return nil, err
	}

	sessions, err := s.sessionUc.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rpcSessions := make([]*protogo.SessionReference, 0, len(sessions))
	for _, session := range sessions {
		rpcSessions = append(rpcSessions, s.sessionToRpcReference(session))
	}

	return &protogo.SessionList{
		Sessions: rpcSessions,
	}, nil
}

func (s *sessionServiceApi) Logout(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	session, err := s.sessionFromAuth(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.sessionUc.Delete(ctx, session.Id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *sessionServiceApi) GetInfo(_ context.Context, _ *emptypb.Empty) (*protogo.ServiceInfo, error) {
	return &protogo.ServiceInfo{
		Name:    "luna.session",
		Version: s.config.Version,
	}, nil
}

// <editor-fold desc="Adapters">

func (s *sessionServiceApi) sessionToRpc(session *domain.Session) *protogo.SessionData {
	return &protogo.SessionData{
		SessionId:  session.Id,
		Name:       session.Name,
		DeviceType: protogo.DeviceType(session.Type),
		Token:      session.Token,
		ExpiresAt:  timestamppb.New(session.ExpiresAt),
	}
}

func (s *sessionServiceApi) sessionToRpcReference(session *domain.Session) *protogo.SessionReference {
	return &protogo.SessionReference{
		SessionId:  session.Id,
		Name:       session.Name,
		DeviceType: protogo.DeviceType(session.Type),
	}
}

// </editor-fold>

func (s *sessionServiceApi) sessionFromAuth(ctx context.Context) (*domain.Session, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	auth := md.Get("authorization")
	if len(auth) == 0 {
		return nil, status.Error(codes.Unauthenticated, "No token provided")
	}
	token := auth[0]

	session, err := s.sessionUc.GetByToken(ctx, token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unknown session")
	}

	return session, nil
}
