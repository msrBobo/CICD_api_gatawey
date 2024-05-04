package grpc_service_clients

import (
	"dennic_api_gateway/genproto/session_service"
	"google.golang.org/grpc"
)

type SessionServiceI interface {
	SessionService() session.SessionServiceClient
}

type SessionService struct {
	sessionService session.SessionServiceClient
}

func NewSessionService(conn *grpc.ClientConn) *SessionService {
	return &SessionService{
		sessionService: session.NewSessionServiceClient(conn),
	}
}

func (s *SessionService) SessionService() session.SessionServiceClient {
	return s.sessionService
}
