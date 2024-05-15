package model_booking_service

type DoctorNote struct {
	Id            int64  `json:"id"`
	AppointmentId int64  `json:"appointment_id"`
	DoctorId      string `json:"doctor_id"`
	PatientId     string `json:"patient_id"`
	Prescription  string `json:"prescription"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type DoctorNotesType struct {
	Count       int64         `json:"count"`
	DoctorNotes []*DoctorNote `json:"doctor_notes"`
}

type CreateDoctorNotesReq struct {
	AppointmentId int64  `json:"appointment_id"`
	DoctorId      string `json:"doctor_id"`
	PatientId     string `json:"patient_id"`
	Prescription  string `json:"prescription"`
}

type UpdateDoctorNoteReq struct {
	DoctorNotesId string `json:"doctor_notes_id"`
	AppointmentId int64  `json:"appointment_id"`
	DoctorId      string `json:"doctor_id"`
	PatientId     string `json:"patient_id"`
	Prescription  string `json:"prescription"`
}
