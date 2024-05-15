package model_user_service

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type User struct {
	Id           string `json:"-"`
	UserOrder    string `json:"-"`
	FirstName    string `json:"first_name" example:"Ali"`
	LastName     string `json:"last_name" example:"Jo'raxonov"`
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
	Id          string `json:"-"`
	FirstName   string `json:"first_name" example:"Ali"`
	LastName    string `json:"last_name" example:"Jo'raxonov'"`
	BrithDate   string `json:"birth_date" example:"2000-01-01"`
	PhoneNumber string `json:"phone_number" example:"+998950230605"`
	Password    string `json:"password" example:"password"`
	Gender      string `json:"gender" example:"male"`
	Code        int64  `json:"-"`
}

type Redis struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name" example:"Ali"`
	LastName    string `json:"last_name" example:"Jo'raxonov"`
	BrithDate   string `json:"birth_date" example:"2000-01-01"`
	PhoneNumber string `json:"phone_number" example:"+998950230605"`
	Password    string `json:"password" example:"password"`
	Gender      string `json:"gender" example:"male"`
	Code        int64  `json:"code"`
}

type MessageRes struct {
	Message string `json:"message"`
}

type VerifyOtpCodeReq struct {
	PhoneNumber string `json:"phone_number" example:"+998950230605"`
	Code        int64  `json:"code" example:"7777"`
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
	Id           string `json:"id"`
	FirstName    string `json:"first_name" `
	LastName     string `json:"last_name" `
	BrithDate    string `json:"birth_date" `
	PhoneNumber  string `json:"phone_number" `
	Gender       string `json:"gender"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type PhoneNumberReq struct {
	PhoneNumber string `json:"phone_number" example:"+998950230605"`
}

type UpdateRefreshTokenUserResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" example:"RefreshToken"`
}

// User info Validate
func (u *Redis) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.PhoneNumber, validation.Required, validation.Length(13, 13), validation.Match(regexp.MustCompile("^\\+[0-9]")).Error("Phone number is not valid")),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 32), validation.Match(regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*()-_=+]")).Error("Password is not valid")),
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 50), validation.Match(regexp.MustCompile("^[A-Z][a-zA-Z']*([\\\\s-][A-Z][a-zA-Z']*)*$")).Error("First name is not valid")),
		validation.Field(&u.LastName, validation.Required, validation.Length(3, 50), validation.Match(regexp.MustCompile("^[A-Z][a-zA-Z']*([\\\\s-][A-Z][a-zA-Z']*)*$")).Error("Last name is not valid")),
	)
}
