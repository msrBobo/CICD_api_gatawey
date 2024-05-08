package model_user_service

type User struct {
	Id           string `json:"-"`
	UserOrder    string `json:"-"`
	FirstName    string `json:"first_name" example:"Ali"`
	LastName     string `json:"last_name" example:"Jo'raxonov'"`
	BrithDate    string `json:"birth_date" example:"2000-01-01"`
	PhoneNumber  string `json:"phone_number" example:"+998950230605"`
	Password     string `json:"password" example:"password"`
	Gender       string `json:"gender" example:"male"`
	RefreshToken string `json:"-"`
}

type Users struct {
	Count uint64  `json:"count"`
	Users []*User `json:"users"`
}

type RegisterRequest struct {
	Id           string `json:"-"`
	FirstName    string `json:"first_name" example:"Ali"`
	LastName     string `json:"last_name" example:"Jo'raxonov'"`
	BrithDate    string `json:"birth_date" example:"2000-01-01"`
	PhoneNumber  string `json:"phone_number" example:"+998950230605"`
	Password     string `json:"password" example:"password"`
	Gender       string `json:"gender" example:"male"`
	RefreshToken string `json:"-"`
	Code         int64  `json:"-"`
}

type Redis struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name" example:"Ali"`
	LastName     string `json:"last_name" example:"Jo'raxonov'"`
	BrithDate    string `json:"birth_date" example:"2000-01-01"`
	PhoneNumber  string `json:"phone_number" example:"+998950230605"`
	Password     string `json:"password" example:"password"`
	Gender       string `json:"gender" example:"male"`
	RefreshToken string `json:"refresh_token"`
	Code         int64  `json:"code"`
}

type MessageRes struct {
	Message string `json:"message"`
}

type ForgetPasswordVerify struct {
	PhoneNumber string `json:"phone_number" example:"+998950230605"`
	Code        int    `json:"code" example:"7777"`
	NewPassword string `json:"new_password" example:"new_password"`
}

type Verify struct {
	PhoneNumber  string `json:"phone_number" example:"+998950230605"`
	Code         int64  `json:"code" example:"7777"`
	PlatformName string `json:"platform_name"`
	PlatformType string `json:"platform_type" example:"mobile"`
	FcmToken     string `json:"fcm_token"`
}

type LoginReq struct {
	PhoneNumber  string `json:"phone_number" example:"+998950230605"`
	Password     string `json:"password" example:"password"`
	PlatformName string `json:"platform_name" `
	PlatformType string `json:"platform_type" example:"mobile"`
	FcmToken     string `json:"fcm_token"`
}

type Response struct {
	FirstName   string `json:"first_name" `
	LastName    string `json:"last_name" `
	BrithDate   string `json:"birth_date" `
	PhoneNumber string `json:"phone_number" `
	Gender      string `json:"gender"`
	AccessToken string `json:"access_token"`
}

type PhoneNumberReq struct {
	PhoneNumber string `json:"phone_number" example:"+998950230605"`
}
