package handler

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

func HandleError(w http.ResponseWriter, err error) {
	errType := reflect.TypeOf(err)
	log.Println(errType)

	errMessage := err.Error()
	code := http.StatusInternalServerError

	if strings.Contains(errMessage, "no document") {
		code = http.StatusNotFound
	}

	if strings.Contains(errMessage, "validator") {
		code = http.StatusBadRequest
	}

	if strings.Contains(errMessage, "token") {
		code = http.StatusUnauthorized
	}

	if strings.Contains(errMessage, "duplicate") {
		code = http.StatusConflict
	}

	if err != nil {
		SendError(w, err, code)
	}
}
