package models

// Response struct
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error         string      `json:"error"`
	Error_message interface{} `json:"error_message"`
	Error_code    string      `json:"error_code"`
}
