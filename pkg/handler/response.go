package handler

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendSuccess(w http.ResponseWriter, data any) {
	resp := &response{
		Success: true,
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func SendError(w http.ResponseWriter, err error, code int) {
	resp := &response{
		Success: false,
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}
