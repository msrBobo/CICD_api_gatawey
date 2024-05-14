package model_healthcare_service

type DoctorServicesReq struct {
	DoctorId         string  `json:"doctor_id" example:"123e4567-e89b-12d3-a456-426614274001"`
	SpecializationId string  `json:"specialization_id" example:"123e4567-e89b-12d3-a456-426614375001"`
	OnlinePrice      float32 `json:"online_price" example:"1.1"`
	OfflinePrice     float32 `json:"offline_price" example:"1.1"`
	Name             string  `json:"name" example:"name"`
	Duration         string  `json:"duration" example:"12:12"`
}

type DoctorServicesRes struct {
	Id               string  `json:"id"`
	Order            int32   `json:"order"`
	DoctorId         string  `json:"doctor_id"`
	SpecializationId string  `json:"specialization_id"`
	OnlinePrice      float32 `json:"online_price"`
	OfflinePrice     float32 `json:"offline_price"`
	Name             string  `json:"name"`
	Duration         string  `json:"duration"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

type ListDoctorServices struct {
<<<<<<< HEAD
=======
	DoctorServices []*DoctorServicesRes `json:"doctor_services"`
>>>>>>> eca93a4806164b242209e3dd00f4c94991ddf34d
	Count          int32                `json:"count"`
	DoctorServices []*DoctorServicesRes `json:"doctorServices"`
}
