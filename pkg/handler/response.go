package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gprestore/gprestore-core/pkg/variable"
)

type response struct {
	Success bool    `json:"success"`
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    any     `json:"data"`
	Next    *string `json:"_next"`
}

func SendSuccess(w http.ResponseWriter, r *http.Request, data any) {
	accessToken, _ := r.Context().Value(variable.ContextKeyAccessToken).(string)
	resp := &response{
		Success: true,
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
		Next:    &accessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func SendError(w http.ResponseWriter, r *http.Request, err error, code int) {
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
