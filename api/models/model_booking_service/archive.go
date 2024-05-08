package model_booking_service

type Archive struct {
	Id                   int64   `json:"id"`
	DoctorAvailabilityId int64   `json:"doctor_availability_id"`
	StartTime            string  `json:"start_time"`
	EndTime              string  `json:"end_time"`
	PatientProblem       string  `json:"patient_problem"`
	Status               string  `json:"status"`
	PaymentType          string  `json:"payment_type"`
	PaymentAmount        float64 `json:"payment_amount"`
}

type ArchivesType struct {
	Count    int64      `json:"count"`
	Archives []*Archive `json:"archives"`
}

type CreateArchiveReq struct {
	DoctorAvailabilityId int64   `json:"doctor_availability_id"`
	StartTime            string  `json:"start_time"`
	EndTime              string  `json:"end_time"`
	PatientProblem       string  `json:"patient_problem"`
	Status               string  `json:"status"`
	PaymentType          string  `json:"payment_type"`
	PaymentAmount        float64 `json:"payment_amount"`
}

type UpdateArchiveReq struct {
	Field                string  `json:"field"`
	Value                string  `json:"value"`
	DoctorAvailabilityId int64   `json:"doctor_availability_id"`
	StartTime            string  `json:"start_time"`
	EndTime              string  `json:"end_time"`
	PatientProblem       string  `json:"patient_problem"`
	Status               string  `json:"status"`
	PaymentType          string  `json:"payment_type"`
	PaymentAmount        float64 `json:"payment_amount"`
}
