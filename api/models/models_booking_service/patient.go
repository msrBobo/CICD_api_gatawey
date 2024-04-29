package models_booking_service

type Patient struct {
	ID             string `json:"id"`
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
	DeletedAt      string `json:"deleted_at"`
}

type Patients struct {
	Count    int64      `json:"count"`
	Patients []*Patient `json:"patients"`
}

type CreatePatientReq struct {
	ID             string `json:"id"`
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
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	BloodGroup     string `json:"blood_group"`
	City           string `json:"city"`
	Country        string `json:"country"`
	PatientProblem string `json:"patient_problem"`
}

type UpdatePhoneNumber struct {
	Field       string `json:"field"`
	Value       string `json:"value"`
	PhoneNumber string `json:"phone_number"`
}
