package model_healthcare_service

type DepartmentReq struct {
	Id               string `json:"id"`
	Name             string `json:"name" example:"name"`
	Description      string `json:"description" example:"description"`
	ImageUrl         string `json:"image_url" example:"http://example.com/image.png"`
	FloorNumber      int32  `json:"floor_number" example:"2"`
	ShortDescription string `json:"short_description" example:"short_description"`
}

type DepartmentRes struct {
	Id               string `json:"id"`
	Order            int32  `json:"order"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ImageUrl         string `json:"image_url"`
	FloorNumber      int32  `json:"floor_number"`
	ShortDescription string `json:"short_description"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type ListDepartments struct {
	Count       int32            `json:"count"`
	Departments []*DepartmentRes `json:"departments"`
}
