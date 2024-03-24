package custom_errors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func NewErrorResponse(status int, message string, errors []string) ErrorResponse {
	return ErrorResponse{
		Status:  status,
		Message: message,
		Errors:  errors,
	}
}

func ReturnHTTPError(w http.ResponseWriter, messages []string) {
	response := ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
		Errors:  messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating error response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonResponse)
}

func ReturnBadRequestError(w http.ResponseWriter, messages []string) {
	response := ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
		Errors:  messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating error response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResponse)
}

func ReturnUnauthorizedError(w http.ResponseWriter, messages []string) {
	response := ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Errors:  messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating error response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(jsonResponse)
}

func ReturnNotFoundError(w http.ResponseWriter, messages []string) {
	response := ErrorResponse{
		Status:  http.StatusNotFound,
		Message: "Not Found",
		Errors:  messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating error response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResponse)
}

func ReturnConflictError(w http.ResponseWriter, messages []string) {
	response := ErrorResponse{
		Status:  http.StatusConflict,
		Message: "Conflict",
		Errors:  messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating error response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusConflict)
	w.Write(jsonResponse)
}

func ReturnForbiddenError(w http.ResponseWriter, messages []string) {
	response := ErrorResponse{
		Status:  http.StatusForbidden,
		Message: "Forbidden",
		Errors:  messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating error response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write(jsonResponse)
}

func ReturnUnprocessableEntityError(w http.ResponseWriter, messages []string) {
	response := ErrorResponse{
		Status:  http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity",
		Errors:  messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating error response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(jsonResponse)
}

func ReturnMethodNotAllowedError(w http.ResponseWriter, messages []string) {
	response := ErrorResponse{
		Status:  http.StatusMethodNotAllowed,
		Message: "Method Not Allowed",
		Errors:  messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating error response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(jsonResponse)
}
