package utils

import "generate-auth-workshop/model"

func CreateErrorResponse(code int, message string) (int, map[string]model.ErrorResponse) {
	errorResponse := map[string]model.ErrorResponse{
		"error": {
			Status: code,
			Message: message,
		},
	}
	return code, errorResponse
}
