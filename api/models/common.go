package models

type FieldValueReq struct {
	Field        string `json:"field"`
	Value        string `json:"value"`
	DeleteStatus bool   `json:"delete_status"`
}

type ListReq struct {
	Page         uint64 `json:"page"`
	Limit        uint64 `json:"limit"`
	OrderBy      string `json:"order_by"`
	Field        string `json:"field"`
	Value        string `json:"value"`
	DeleteStatus bool   `json:"delete_status"`
}

type StatusRes struct {
	Status bool `json:"status"`
}

type AccessToken struct {
	Token string `json:"token"`
}
