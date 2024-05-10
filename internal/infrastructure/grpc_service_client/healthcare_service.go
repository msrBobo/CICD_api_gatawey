package grpc_service_clients

import (
	"dennic_api_gateway/genproto/healthcare-service"

	"google.golang.org/grpc"
)

type HealthcareServiceI interface {
	DepartmentService() healthcare.DepartmentServiceClient
	DoctorService() healthcare.DoctorServiceClient
	DoctorsService() healthcare.DoctorsServiceClient
	DoctorWorkingHoursService() healthcare.DoctorWorkingHoursServiceClient
	SpecializationService() healthcare.SpecializationServiceClient
	ReasonsService() healthcare.ReasonsServiceClient
}

type HealthcareService struct {
	departmentService         healthcare.DepartmentServiceClient
	doctorService             healthcare.DoctorServiceClient
	doctorsService            healthcare.DoctorsServiceClient
	doctorWorkingHoursService healthcare.DoctorWorkingHoursServiceClient
	specializationService     healthcare.SpecializationServiceClient
	reasonsService            healthcare.ReasonsServiceClient
}

func NewHealthcareService(conn *grpc.ClientConn) *HealthcareService {
	return &HealthcareService{
		departmentService:         healthcare.NewDepartmentServiceClient(conn),
		doctorService:             healthcare.NewDoctorServiceClient(conn),
		doctorsService:            healthcare.NewDoctorsServiceClient(conn),
		doctorWorkingHoursService: healthcare.NewDoctorWorkingHoursServiceClient(conn),
		specializationService:     healthcare.NewSpecializationServiceClient(conn),
		reasonsService:            healthcare.NewReasonsServiceClient(conn),
	}
}

func (s *HealthcareService) DepartmentService() healthcare.DepartmentServiceClient {
	return s.departmentService
}

func (s *HealthcareService) DoctorService() healthcare.DoctorServiceClient {
	return s.doctorService
}

func (s *HealthcareService) DoctorsService() healthcare.DoctorsServiceClient {
	return s.doctorsService
}

func (s *HealthcareService) DoctorWorkingHoursService() healthcare.DoctorWorkingHoursServiceClient {
	return s.doctorWorkingHoursService
}

func (s *HealthcareService) SpecializationService() healthcare.SpecializationServiceClient {
	return s.specializationService
}

func (s *HealthcareService) ReasonsService() healthcare.ReasonsServiceClient {
	return s.reasonsService
}
