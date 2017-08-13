package handler

import (
	"net/http"
	"encoding/json"
	"zipper/utils"
	"fmt"
	"github.com/gorilla/mux"
	"zipper/controller/auth"
	"zipper/controller/response"
	"zipper/controller/model"
	"zipper/files"
	"zipper/log"
)

func DownloadZip(w http.ResponseWriter, r *http.Request) {
	queryInfo := extractQueryInfo(r)

	if err := auth.ZipGetValidator.Validate(queryInfo); err != nil {
		response.Unauthorized(w, err)
		return
	}

	info := fetchZipInfo(w, queryInfo.ZipKey)
	if info == nil {
		return
	}

	fileInfos := decorateFileNamesInZip(info.FileInfos, queryInfo.ZipName)
	response.ZipStreamAsResponse(w, fileInfos, queryInfo.ZipName)
}

func extractQueryInfo(r *http.Request) *model.QueryInfo {
	vars := mux.Vars(r)
	queryParams := r.URL.Query()
	urlPath := r.URL.Path

	return &model.QueryInfo{
		ZipKey: vars["key"],
		Timestamp: queryParams.Get("timestamp"),
		Token: queryParams.Get("token"),
		ZipName: queryParams.Get("name"),
		URI: urlPath,
	}
}

func fetchZipInfo(w http.ResponseWriter, key string) *ZipInfo {
	bytes, err := utils.GetRedisClient().GetValue(key)
	if err != nil {
		response.NotFound(w)
		return nil
	}

	zipInfo := &ZipInfo{}
	if err1 := json.Unmarshal(bytes, zipInfo); err1 != nil {
		msg := "Unmarshal zip info failed"
		log.Error(msg, key, string(bytes))
		response.ServerError(w, msg)
		return nil
	}
	return zipInfo
}

func decorateFileNamesInZip(fileInfos []files.FileInfo, zipName string) []files.FileInfo {
	filenameMap := make(map[string]int)
	topFolderName, _ := utils.FileNameAndExt(zipName)

	decorated := []files.FileInfo{}
	for _, fileInfo := range fileInfos {
		name := preventDuplicateFilename(&filenameMap, fileInfo.Filename)
		fileInfo.Filename = addTopFolderInPath(name, topFolderName)
		decorated = append(decorated, fileInfo)
	}
	return decorated
}

func addTopFolderInPath(filePath string, topFolderName string) string {
	return fmt.Sprintf("%s/%s", topFolderName, filePath)
}

func preventDuplicateFilename(filenameMap *map[string]int, filename string) string {
	existingNameCount, isKeyExisting := (*filenameMap)[filename]
	if !isKeyExisting {
		(*filenameMap)[filename] = 1
		return filename
	}

	(*filenameMap)[filename] = existingNameCount + 1

	name, ext := utils.FileNameAndExt(filename)
	return fmt.Sprintf("%s_%d%s", name, existingNameCount, ext)
}
