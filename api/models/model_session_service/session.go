package model_session_service

type SessionReq struct {
	IpAddress    string `json:"ip_address"`
	UserId       string `json:"user_id"`
	FcmToken     string `json:"fcm_token"`
	PlatformName string `json:"platform_name"`
	PlatformType string `json:"platform_type"`
}

type SessionRes struct {
	Id           string `json:"id"`
	Order        int32  `json:"order"`
	IpAddress    string `json:"ip_address"`
	UserId       string `json:"user_id"`
	FcmToken     string `json:"fcm_token"`
	PlatformName string `json:"platform_name"`
	PlatformType string `json:"platform_type"`
	LoginAt      string `json:"login_at"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type ListSessions struct {
	Sessions []*SessionRes `json:"sessions"`
	Count    int32         `json:"count"`
}
