package entity

import "time"

type DoctorNote struct {
	Id           int64     `json:"id"`
	AppointmetId int64     `json:"appointment_id"`
	DoctorId     string    `json:"doctor_id"`
	PatientId    string    `json:"patient_id"`
	Prescription string    `json:"prescription"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type DoctorNotes struct {
	Count       int64         `json:"count"`
	DoctorNotes []*DoctorNote `json:"doctor_notes"`
}

type CreateDoctorNoteReq struct {
	AppointmetId int64  `json:"appointment_id"`
	DoctorId     string `json:"doctor_id"`
	PatientId    string `json:"patient_id"`
	Prescription string `json:"prescription"`
}

type UpdateDoctorNoteReq struct {
	Field         string `json:"field"`
	Value         string `json:"value"`
	AppointmentId int64  `json:"appointment_id"`
	DoctorId      string `json:"doctor_id"`
	PatientId     string `json:"patient_id"`
	Prescription  string `json:"prescription"`
}
