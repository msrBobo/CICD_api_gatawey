package model_common

// ResponseError ...
type ResponseError struct {
	Code    string `json:"status"`
	Message string `json:"message"`
	Date    string `json:"date"`
}

// StandardErrorModel ...
type StandardErrorModel struct {
	Error ResponseError `json:"error"`
}
