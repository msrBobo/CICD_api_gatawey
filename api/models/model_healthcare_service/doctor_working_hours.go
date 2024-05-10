package model_healthcare_service

type DoctorWorkingHoursRes struct {
	Id         int32  `json:"id"`
	DoctorId   string `json:"doctorId"`
	DayOfWeek  string `json:"dayOfWeek"`
	StartTime  string `json:"startTime"`
	FinishTime string `json:"finishTime"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type DoctorWorkingHoursReq struct {
	DoctorId   string `json:"doctorId"`
	DayOfWeek  string `json:"dayOfWeek"`
	StartTime  string `json:"startTime"`
	FinishTime string `json:"finishTime"`
}

type ListDoctorWorkingHours struct {
	ListDWH []*DoctorWorkingHoursRes `json:"doctor_working_hours"`
	Count   int32                    `json:"count"`
}
