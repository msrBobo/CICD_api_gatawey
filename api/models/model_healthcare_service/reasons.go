package model_healthcare_service

type ReasonsReq struct {
	Name             string `json:"name"`
	SpecializationId string `json:"specializationId"`
	ImageUrl         string `json:"imageUrl"`
}
type ReasonsRes struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	SpecializationId string `json:"specializationId"`
	ImageUrl         string `json:"imageUrl"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

type ListReasons struct {
	Reasons []*ReasonsRes `json:"reasons"`
	Count   int32         `json:"count"`
}
