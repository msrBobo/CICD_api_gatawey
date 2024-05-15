package model_booking_service

type Patient struct {
	Id             string `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	BloodGroup     string `json:"blood_group"`
	PhoneNumber    string `json:"phone_number"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PatientProblem string `json:"patient_problem"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type PatientsType struct {
	Count    int64      `json:"count"`
	Patients []*Patient `json:"patients"`
}

type CreatePatientReq struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	BloodGroup     string `json:"blood_group"`
	PhoneNumber    string `json:"phone_number"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PatientProblem string `json:"patient_problem"`
}

type UpdatePatientReq struct {
	PatientId      string `json:"patient_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	BloodGroup     string `json:"blood_group"`
	PhoneNumber    string `json:"phone_number"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PatientProblem string `json:"patient_problem"`
}
