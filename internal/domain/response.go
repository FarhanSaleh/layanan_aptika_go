package domain

type DefaultResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorValidationResponse struct {
	Message string `json:"message"`
	Errors  []ErrorsValidation
}

type ErrorsValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}