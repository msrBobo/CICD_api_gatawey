package model_healthcare_service

type DoctorServices struct {
	Id                 string  `json:"-"`
	DoctorServiceOrder int32   `json:"-"`
	DoctorId           string  `json:"-"`
	Specialization     string  `json:"specialization"`
	OnlinePrice        float64 `json:"online_price"`
	OfflinePrice       float64 `json:"offline_price"`
	Name               string  `json:"name"`
	Duration           string  `json:"duration"`
}

type DoctorServicesType struct {
	DoctorServices []*DoctorServices
	Count          int32
}
