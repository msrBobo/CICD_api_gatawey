package entity

import (
	"time"

	"github.com/rickb777/date"
)

type Doctor struct {
	Id           int64     `json:"id"`
	DepartmentId int64     `json:"department_id"`
	DoctorId     int64     `json:"doctor_id"`
	DoctorDate   date.Date `json:"doctor_date"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Status       bool      `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type DoctorTimes struct {
	Count  int64     `json:"count"`
	Doctor []*Doctor `json:"doctor"`
}

type CreateDoctorTimeReq struct {
	DepartmentId int64     `json:"department_id"`
	DoctorId     int64     `json:"doctor_id"`
	DoctorDate   date.Date `json:"doctor_date"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Status       bool      `json:"status"`
}

type UpdateDoctorTimeReq struct {
	Field        string    `json:"field"`
	Value        string    `json:"value"`
	DepartmentId int64     `json:"department_id"`
	DoctorId     int64     `json:"doctor_id"`
	DoctorDate   date.Date `json:"doctor_date"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Status       bool      `json:"status"`
}
