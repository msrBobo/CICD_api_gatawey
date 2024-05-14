package model_healthcare_service

type DoctorServicesReq struct {
	DoctorId         string  `json:"doctorId"`
	SpecializationId string  `json:"specializationId"`
	OnlinePrice      float32 `json:"onlinePrice"`
	OfflinePrice     float32 `json:"offlinePrice"`
	Name             string  `json:"name"`
	Duration         string  `json:"duration"`
}

type DoctorServicesRes struct {
	Id               string  `json:"id"`
	Order            int32   `json:"order"`
	DoctorId         string  `json:"doctorId"`
	SpecializationId string  `json:"specializationId"`
	OnlinePrice      float32 `json:"onlinePrice"`
	OfflinePrice     float32 `json:"offlinePrice"`
	Name             string  `json:"name"`
	Duration         string  `json:"duration"`
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
}

type ListDoctorServices struct {
	Count          int32                `json:"count"`
	DoctorServices []*DoctorServicesRes `json:"doctorServices"`
}
