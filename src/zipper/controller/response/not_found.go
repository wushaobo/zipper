package response

import (
	"net/http"
	"encoding/json"
	"zipper/files"
)

func NotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

func URLsNotFound(w http.ResponseWriter, urlsNotFound []files.FileInfo) {
	NotFound(w)

	type Response struct {
		Message string `json:"message"`
		URLsNotFound []files.FileInfo `json:"urls_not_found"`
	}

	marshaled, _ := json.Marshal(&Response{
		Message: "urls not found",
		URLsNotFound: urlsNotFound,
	})
	w.Write(marshaled)
}