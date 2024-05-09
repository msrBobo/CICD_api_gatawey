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
	Name         string `json:"name"`
	Description  string `json:"description"`
	DepartmentId string `json:"department_id"`
	ImageUrl     string `json:"image_url"`
}

type ListSpecializations struct {
	Specializations []*SpecializationRes `json:"specializations"`
	Count           int32                `json:"count"`
}
