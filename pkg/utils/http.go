package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dto interface{}) *APIErrorResponse {
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		validationErr := NewAPIErrorResponse()
		validationErr.AddGeneralError("Invalid JSON data provided")
		w.WriteHeader(validationErr.Code)
		json.NewEncoder(w).Encode(validationErr)
		return validationErr
	}
	return nil
}
