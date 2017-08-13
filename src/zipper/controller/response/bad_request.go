package response

import (
	"net/http"
	"encoding/json"
)

func badRequest(w http.ResponseWriter, marshaledJson []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(marshaledJson)
}

func BadRequestWithMessage(w http.ResponseWriter, message string) {
	type Response struct {
		Message string		`json:"message"`
	}
	marshaled, _ := json.Marshal(&Response{
		Message: message,
	})
	badRequest(w, marshaled)
}
