package utils

import (
	"encoding/json"
	"net/http"
)

type FieldError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type APIErrorResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors,omitempty"`
}

func (e *APIErrorResponse) Error() string {
	errJSON, _ := json.Marshal(e)
	return string(errJSON)
}

func NewAPIErrorResponse(code ...int) *APIErrorResponse {
	errorCode := http.StatusBadRequest
	if len(code) > 0 {
		errorCode = code[0]
	}
	return &APIErrorResponse{
		Code:   errorCode,
		Errors: []FieldError{},
	}
}

func (e *APIErrorResponse) AddFieldError(field, message string) {
	e.Errors = append(e.Errors, FieldError{
		Field:   field,
		Message: message,
	})
}

func (e *APIErrorResponse) AddGeneralError(message string) {
	e.Message = message
}

func NewInternalServerError(message ...string) *APIErrorResponse {
	finalMessage := "An unexpected error occurred. Please try again later."
	if len(message) > 0 {
		finalMessage = message[0]
	}
	return &APIErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: finalMessage,
	}
}

func (e *APIErrorResponse) HasErrors() bool {
	return len(e.Errors) > 0 || e.Message != ""
}
