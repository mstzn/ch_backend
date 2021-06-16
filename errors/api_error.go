package errors

type ApiError struct {
	Message string `json:"message"`
	Code string `json:"code"`
}