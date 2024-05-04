package entity

type FieldValueReq struct {
	Field    string `json:"field"`
	Value    string `json:"value"`
	IsActive bool   `json:"is_active"`
}

type DeleteStatus struct {
	Status bool `json:"status"`
}

type GetAllRequest struct {
	Field    string `json:"field"`
	Value    string `json:"value"`
	IsActive bool   `json:"is_active"`
	Page     uint64 `json:"page"`
	Limit    uint64 `json:"limit"`
	OrderBy  string `json:"order_by"`
	Search   string `json:"search"`
}
