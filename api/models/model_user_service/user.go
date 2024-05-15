package model_user_service

type GetUserResp struct {
	Id          string `json:"id"`
	UserOrder   uint64 `json:"user_order"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BrithDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	ImageUrl    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListUserResp struct {
	Count uint64        `json:"count"`
	Users []GetUserResp `json:"users"`
}

type UpdUserReq struct {
	Id        string `json:"id" example:"UUID"`
	FirstName string `json:"first_name" example:"To'rahon"`
	LastName  string `json:"last_name" example:"To'rayevich"`
	BrithDate string `json:"birth_date" example:"1990-01-01"`
	Gender    string `json:"gender" example:"male"`
	ImageUrl  string `json:"image_url"`
}

type UpdUserResp struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BrithDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	ImageUrl  string `json:"image_url"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllReq struct {
	Page    uint64 `json:"page"`
	Limit   uint64 `json:"limit"`
	Field   string `json:"field"`
	Value   string `json:"value"`
	OrderBy string `json:"order_by"`
}

type GetUserReq struct {
	Field    string `json:"field"`
	Value    string `json:"value"`
	IsActive bool   `json:"is_active"`
}

type DeleteUserReq struct {
	Field        string `json:"field"`
	Value        string `json:"value"`
	DeleteStatus bool   `json:"delete_status"`
}

type CheckUserFieldResp struct {
	Status bool `json:"status"`
}
