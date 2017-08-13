package files

import (
	"archive/zip"
	"net/http"
	"io"
	"time"
	"zipper/log"
)

const (
	HTTP_GET_TIMEOUT = 7200 * time.Second
)

func RealWriteFilesToZip(w http.ResponseWriter, srcFiles []FileInfo) {
	zipper := zip.NewWriter(w)
	defer zipper.Close()

	writeToZipWithFunc(zipper, srcFiles, realWriteFileToZip)
}

func writeToZipWithFunc(zipper *zip.Writer, srcFiles []FileInfo, writeFunc func(*zip.Writer, FileInfo, time.Time)) {
	unifiedModTime := time.Now().UTC()
	for _, fileInfo := range srcFiles {
		if fileInfo.IsFolder {
			zipper.Create(fileInfo.Filename)
			continue
		}
		writeFunc(zipper, fileInfo, unifiedModTime)
	}
}

func realWriteFileToZip(zipper *zip.Writer, fileInfo FileInfo, modTime time.Time) {
	log.Info("Start writing resource: ", fileInfo.URL)

	srcURL := fileInfo.URL
	resp, err := httpGet(srcURL)

	if err != nil{
		log.Error("Failed to fetch file. error: ", err.Error())
		return
	} else if resp != nil && resp.StatusCode != http.StatusOK {
		log.Error("Failed to fetch file. ", "status code: ", resp.StatusCode, "url: ", srcURL)
		return
	}

	defer resp.Body.Close()
	written, writeErr := writeContentFromReader(zipper, resp.Body, fileInfo.Filename, modTime)
	if writeErr == nil {
		log.Info("Written length: ", written)
	} else {
		log.Error("Written error: ", writeErr.Error())
	}
}

func writeContentFromReader(zipper *zip.Writer, reader io.ReadCloser, filename string, modTime time.Time) (written int64, err error) {
	fileHeader := buildFileHeader(filename, modTime)
	writer, _ := zipper.CreateHeader(&fileHeader)

	return io.CopyBuffer(writer, reader, make([]byte, 32*1024))
}

func buildFileHeader(filename string, modTime time.Time) zip.FileHeader {
	GENERAL_PURPOSE_BIT_FLAG_TO_ENABLE_UNICODE_FOR_FILENAME := uint16(0x800)

	header := zip.FileHeader{
		Name: filename,
		Method: zip.Store,
		Flags: GENERAL_PURPOSE_BIT_FLAG_TO_ENABLE_UNICODE_FOR_FILENAME,
	}
	header.SetModTime(modTime)
	return header
}

func httpGet(url string) (*http.Response, error) {
	timeout := time.Duration(HTTP_GET_TIMEOUT)
	client := http.Client{
	    Timeout: timeout,
	}
	return client.Get(url)
}
