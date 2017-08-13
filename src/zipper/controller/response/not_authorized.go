package response

import (
	"net/http"
	"encoding/json"
)

func Unauthorized(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(marshalJson(err.Error()))
}

func marshalJson(message string) []byte {
	type Response struct {
		Message string		`json:"message"`
	}
	marshaled, _ := json.Marshal(&Response{
		Message: message,
	})
	return marshaled
}