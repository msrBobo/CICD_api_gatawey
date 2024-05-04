package grpc_service_clients

import (
	"dennic_api_gateway/genproto/booking_service"

	"google.golang.org/grpc"
)

type BookingServiceI interface {
	PatientService() booking_service.PatientsServiceClient
	ArchiveService() booking_service.ArchiveServiceClient
	BookedAppointment() booking_service.BookedAppointmentsServiceClient
	DoctorTimes() booking_service.DoctorTimeServiceClient
	DoctorNotes() booking_service.DoctorNotesServiceClient
}

type BookingService struct {
	patientService    booking_service.PatientsServiceClient
	archiveService    booking_service.ArchiveServiceClient
	bookedAppointment booking_service.BookedAppointmentsServiceClient
	doctorTimes       booking_service.DoctorTimeServiceClient
	doctorNotes       booking_service.DoctorNotesServiceClient
}

func NewBookingService(conn *grpc.ClientConn) *BookingService {
	return &BookingService{
		patientService:    booking_service.NewPatientsServiceClient(conn),
		archiveService:    booking_service.NewArchiveServiceClient(conn),
		bookedAppointment: booking_service.NewBookedAppointmentsServiceClient(conn),
		doctorTimes:       booking_service.NewDoctorTimeServiceClient(conn),
		doctorNotes:       booking_service.NewDoctorNotesServiceClient(conn),
	}
}

func (s *BookingService) PatientService() booking_service.PatientsServiceClient {
	return s.patientService
}

func (s *BookingService) ArchiveService() booking_service.ArchiveServiceClient {
	return s.archiveService
}

func (s *BookingService) BookedAppointment() booking_service.BookedAppointmentsServiceClient {
	return s.bookedAppointment
}

func (s *BookingService) DoctorTimes() booking_service.DoctorTimeServiceClient {
	return s.doctorTimes
}

func (s *BookingService) DoctorNotes() booking_service.DoctorNotesServiceClient {
	return s.doctorNotes
}
