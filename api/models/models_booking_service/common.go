package models_booking_service

type FieldValueReq struct {
	Field    string `json:"field"`
	Value    string `json:"value"`
	IsActive bool   `json:"is_active"`
}

type DeleteStatus struct {
	Status bool `json:"status"`
}

type GetAllRequest struct {
	Field    string `json:"field" example:"first_name"`
	Value    string `json:"value" example:"A"`
	IsActive bool   `json:"is_active" example:"true"`
	Page     uint64 `json:"page" example:"1"`
	Limit    uint64 `json:"limit" example:"10"`
	OrderBy  string `json:"order_by" example:"last_name"`
	Search   string `json:"search"`
}

type StatusType struct {
	Status bool `json:"status"`
}

type Errors struct {
	StatusCode int   `json:"status_code"`
	ErrorRes   error `json:"error_res"`
}
