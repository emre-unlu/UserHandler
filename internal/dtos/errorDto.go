package dtos

type ErrorDto struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewErrorDto(message string, code int) ErrorDto {
	return ErrorDto{
		Message: message,
		Code:    code,
	}
}
