package model_healthcare_service

type ReasonsReq struct {
	Name             string `json:"name" example:"name"`
	SpecializationId string `json:"specialization_id" example:"123e4567-e89b-12d3-a456-426614375001"`
	ImageUrl         string `json:"image_url" example:"http://example.com/image.png"`
}
type ReasonsRes struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	SpecializationId string `json:"specialization_id"`
	ImageUrl         string `json:"image_url"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type ListReasons struct {
	Reasons []*ReasonsRes `json:"reasons"`
	Count   int32         `json:"count"`
}
