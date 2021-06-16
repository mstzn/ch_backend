package errors

type InvalidRequestError struct {
	ApiError
}

func GetInvalidRequest(message string) InvalidRequestError {
	error := InvalidRequestError{}
	error.Message = message

	return error
}