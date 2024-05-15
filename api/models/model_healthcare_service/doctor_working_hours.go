package model_healthcare_service

type DoctorWorkingHoursRes struct {
	Id         int32  `json:"id"`
	DoctorId   string `json:"doctor_id"`
	DayOfWeek  string `json:"day_of_week"`
	StartTime  string `json:"start_time"`
	FinishTime string `json:"finish_time"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type DoctorWorkingHoursReq struct {
	Id         string `json:"id" example:"123e4567-e89b-12d3-a456-426614274001"`
	DoctorId   string `json:"doctor_id" example:"123e4567-e89b-12d3-a456-426614274001"`
	DayOfWeek  string `json:"day_of_week" example:"Monday"`
	StartTime  string `json:"start_time" example:"12:00:00"`
	FinishTime string `json:"finish_time" example:"12:00:00"`
}

type ListDoctorWorkingHours struct {
	Count   int32                    `json:"count"`
	ListDWH []*DoctorWorkingHoursRes `json:"doctor_working_hours"`
}
