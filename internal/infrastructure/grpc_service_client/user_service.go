package grpc_service_clients

import (
	user_service "dennic_api_gateway/genproto/user_service"

	"google.golang.org/grpc"
)

type UserServiceI interface {
	AdminService() user_service.AdminServiceClient
	UserService() user_service.UserServiceClient
}

type UserService struct {
	adminService user_service.AdminServiceClient
	userService  user_service.UserServiceClient
}

func NewUserService(conn *grpc.ClientConn) *UserService {
	return &UserService{
		adminService: user_service.NewAdminServiceClient(conn),
		userService:  user_service.NewUserServiceClient(conn),
	}
}

func (s *UserService) AdminService() user_service.AdminServiceClient {
	return s.adminService
}

func (s *UserService) UserService() user_service.UserServiceClient {
	return s.userService
}
