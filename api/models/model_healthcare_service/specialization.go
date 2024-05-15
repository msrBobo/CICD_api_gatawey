package model_healthcare_service

type SpecializationRes struct {
	ID           string `json:"id"`
	Order        int32  `json:"order"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DepartmentId string `json:"department_id"`
	ImageUrl     string `json:"image_url"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type SpecializationReq struct {
	Id           string `json:"id" example:"123e4567-e89b-12d3-a456-426614174003"`
	Name         string `json:"name" example:"Specialization"`
	Description  string `json:"description" example:"Specialization description"`
	DepartmentId string `json:"department_id" example:"123e4567-e89b-12d3-a456-426614174003"`
	ImageUrl     string `json:"image_url" example:"http://example.com/image.png"`
}

type ListSpecializations struct {
	Count           int32                `json:"count"`
	Specializations []*SpecializationRes `json:"specializations"`
}
