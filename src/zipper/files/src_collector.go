package files

import (
	"net/http"
	"strconv"
	"fmt"
)

type ResourcesInfo struct {
	FileInfos []FileInfo `json:"file_infos"`
	Token string `json:"token"`
	Timestamp string `json:"timestamp"`
}

type FileInfo struct {
	Filename string `json:"filename"`
	IsFolder bool `json:"is_empty_folder"`
	URL string `json:"url"`
	Size int64 `json:"size,omitempty"`
}

func CollectSrcFileInfos(resourcesInfo *ResourcesInfo) ([]FileInfo, []FileInfo) {
	files := []FileInfo{}
	resourcesNotFound := []FileInfo{}

	for _, fileInfo := range resourcesInfo.FileInfos {
		var fileLength int64
		if !validateResource(fileInfo, &fileLength) {
			resourcesNotFound = append(resourcesNotFound, fileInfo)
			continue
		}

		fileInfo.Size = fileLength
		files = append(files, fileInfo)
	}
	return files, resourcesNotFound
}

func validateResource(fileInfo FileInfo, fileLength *int64) bool {
	if fileInfo.IsFolder {
		return true
	}

	resp, err := http.Get(fileInfo.URL)
	if err != nil {
		fmt.Println("err:")
		fmt.Println(err)
		return false
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp.StatusCode:")
		fmt.Println(resp.StatusCode)
		return false
	}

	defer resp.Body.Close()

	*fileLength, _ = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	return true
}
