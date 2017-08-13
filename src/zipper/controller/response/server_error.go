package response

import (
	"net/http"
	"encoding/json"
)

func ServerError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	type Response struct {
		Message string		`json:"message"`
	}
	marshaled, _ := json.Marshal(&Response{
		Message: message,
	})
	w.Write(marshaled)
}
