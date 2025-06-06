package domain

type DefaultResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorValidationResponse struct {
	Message string `json:"message"`
	Errors  []ErrorsValidation `json:"errors"`
}

type ErrorsValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}