package shared

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	IsSucess bool   `json:"is_success"`       // Indicates if the operation was successful
	Message  string `json:"message"`          // Short message
	Data     any    `json:"data,omitempty"`   // Optional payload
	Errors   any    `json:"errors,omitempty"` // Optional validation or system errors
}

func JSON(w http.ResponseWriter, statusCode int, res BaseResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func Success(w http.ResponseWriter, message string, data any) {
	JSON(w, http.StatusOK, BaseResponse{
		IsSucess: true,
		Message:  message,
		Data:     data,
	})
}

func Error(w http.ResponseWriter, statusCode int, message string, err any) {
	JSON(w, statusCode, BaseResponse{
		IsSucess: false,
		Message:  message,
		Errors:   err,
	})
}
