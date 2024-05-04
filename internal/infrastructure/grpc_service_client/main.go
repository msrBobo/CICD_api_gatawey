package grpc_service_clients

import (
	"dennic_api_gateway/internal/pkg/config"
	"fmt"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	BookingService() BookingServiceI
	HealthcareService() HealthcareServiceI
	UserService() UserServiceI
	SessionService() SessionServiceI
	Close()
}

type serviceClient struct {
	bookingService    BookingServiceI
	healthcareService HealthcareServiceI
	userService       UserServiceI
	sessionService    SessionServiceI
	connections       []*grpc.ClientConn
}

func New(cfg *config.Config) (ServiceClient, error) {
	connBookingService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.BookingService.Host, cfg.BookingService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	connHealthcareService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.HealthcareService.Host, cfg.HealthcareService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserService.Host, cfg.UserService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	connSessionService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.SessionService.Host, cfg.SessionService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceClient{
		bookingService:    NewBookingService(connBookingService),
		healthcareService: NewHealthcareService(connHealthcareService),
		userService:       NewUserService(connUserService),
		sessionService:    NewSessionService(connSessionService),
		connections:       []*grpc.ClientConn{connBookingService, connHealthcareService, connUserService},
	}, nil
}

func (s *serviceClient) BookingService() BookingServiceI {
	return s.bookingService
}

func (s *serviceClient) HealthcareService() HealthcareServiceI {
	return s.healthcareService
}

func (s *serviceClient) UserService() UserServiceI {
	return s.userService
}

func (s *serviceClient) SessionService() SessionServiceI {
	return s.sessionService
}

func (s *serviceClient) Close() {
	for _, conn := range s.connections {
		if err := conn.Close(); err != nil {
			// should be replaced by logger soon
			fmt.Printf("error while closing grpc connection: %v", err)
		}
	}
}
