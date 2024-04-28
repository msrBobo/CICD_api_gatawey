package entity

import (
	"github.com/rickb777/date"
	"time"
)

type Patient struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	BirthDate      date.Date `json:"birth_date"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	BloodGroup     string    `json:"blood_group"`
	PhoneNumber    string    `json:"phone_number"`
	City           string    `json:"city"`
	Country        string    `json:"country"`
	PatientProblem string    `json:"patient_problem"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type Patients struct {
	Count    int64      `json:"count"`
	Patients []*Patient `json:"patients"`
}

type CreatePatientReq struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	BirthDate      date.Date `json:"birth_date"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	BloodGroup     string    `json:"blood_group"`
	PhoneNumber    string    `json:"phone_number"`
	City           string    `json:"city"`
	Country        string    `json:"country"`
	PatientProblem string    `json:"patient_problem"`
}

type UpdatePatientReq struct {
	Field          string    `json:"field"`
	Value          string    `json:"value"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	BirthDate      date.Date `json:"birth_date"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	BloodGroup     string    `json:"blood_group"`
	City           string    `json:"city"`
	Country        string    `json:"country"`
	PatientProblem string    `json:"patient_problem"`
}

type UpdatePhoneNumber struct {
	Field       string `json:"field"`
	Value       string `json:"value"`
	PhoneNumber string `json:"phone_number"`
}
