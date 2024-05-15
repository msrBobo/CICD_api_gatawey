package model_common

// ResponseError ...
type ResponseError struct {
	Code    string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// StandardErrorModel ...
type StandardErrorModel struct {
	Error ResponseError `json:"error"`
}
