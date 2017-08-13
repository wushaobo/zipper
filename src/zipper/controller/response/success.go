package response

import (
	"fmt"
	"net/http"
	"zipper/files"
	"strconv"
	"encoding/json"
	"net/url"
)

func ZipKeyAsResponse(w http.ResponseWriter, zipKey string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	type Response struct {
		ZipKey string `json:"zip_key"`
	}
	marshaled, _ := json.Marshal(&Response{
		ZipKey: zipKey,
	})
	w.Write(marshaled)
}

func ZipStreamAsResponse(w http.ResponseWriter, srcFiles []files.FileInfo, zipName string) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires","0")

	w.Header().Set("Content-Type", "application/zip")
	setFilenameInContentDisposition(w, zipName)
	setContentLength(w, srcFiles)

	w.WriteHeader(http.StatusOK)
	files.RealWriteFilesToZip(w, srcFiles)
}

func setFilenameInContentDisposition(w http.ResponseWriter, filename string) {
	filename_ascii := strconv.QuoteToASCII(filename)
	filename_rfc2231 := encodeRfc2231(filename)
	value := fmt.Sprintf("attachment; filename=%s; filename*=%s", filename_ascii, filename_rfc2231)
	w.Header().Set("Content-Disposition", value)
}

func encodeRfc2231(name string) string {
	return fmt.Sprintf("utf-8''%s", url.QueryEscape(name))
}

func setContentLength(w http.ResponseWriter, srcFiles []files.FileInfo) {
	totalLength := files.FakeWriteFilesToZip(srcFiles)
	w.Header().Set("Content-Length", strconv.FormatInt(int64(totalLength), 10))
}