package entity

import (
	"github.com/rickb777/date"
	"time"
)

type Appointment struct {
	Id              int64     `json:"id"`
	DepartmentId    string    `json:"department_id"`
	DoctorId        string    `json:"doctor_id"`
	PatientId       string    `json:"patient_id"`
	AppointmentDate date.Date `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Duration        int64     `json:"duration"`
	Key             string    `json:"key"`
	ExpiresAt       time.Time `json:"expires_at"`
	PatientStatus   bool      `json:"patient_status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type Appointments struct {
	Count        int64          `json:"count"`
	Appointments []*Appointment `json:"appointments"`
}

type CreateAppointmentReq struct {
	DepartmentId    string    `json:"department_id"`
	DoctorId        string    `json:"doctor_id"`
	PatientId       string    `json:"patient_id"`
	AppointmentDate date.Date `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Duration        int64     `json:"duration"`
	Key             string    `json:"key"`
	ExpiresAt       time.Time `json:"expires_at"`
	PatientStatus   bool      `json:"patient_status"`
}

type UpdateAppointmentReq struct {
	AppointmentDate date.Date `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Duration        int64     `json:"duration"`
	Key             string    `json:"key"`
	ExpiresAt       time.Time `json:"expires_at"`
	PatientStatus   bool      `json:"patient_status"`
	Field           string    `json:"field"`
	Value           string    `json:"value"`
}
