package model_booking_service

type DoctorTime struct {
	Id           int64  `json:"id"`
	DepartmentId string `json:"department_id"`
	DoctorId     string `json:"doctor_id"`
	DoctorDate   string `json:"doctor_date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Status       string `json:"status"`
}

type DoctorTimesType struct {
	Count       int64         `json:"count"`
	DoctorTimes []*DoctorTime `json:"doctor_times"`
}

type CreateDoctorTimeReq struct {
	DepartmentId string `json:"department_id"`
	DoctorId     string `json:"doctor_id"`
	DoctorDate   string `json:"doctor_date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Status       string `json:"status"`
}

type UpdateDoctorTimeReq struct {
	Field        string `json:"field"`
	Value        string `json:"value"`
	DepartmentId string `json:"department_id"`
	DoctorId     string `json:"doctor_id"`
	DoctorDate   string `json:"doctor_date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Status       string `json:"status"`
}
