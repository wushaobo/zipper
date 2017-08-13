package handler

import (
	"net/http"
	"zipper/files"
	"encoding/json"
	"zipper/utils"
	"zipper/controller/response"
	"zipper/controller/auth"
	"zipper/log"
	"time"
)

var (
	StatusDuration = 24 * time.Hour
)

func CreateZipInfo(w http.ResponseWriter, r *http.Request) {
	resourcesInfo := extractInfoFrom(r)

	if resourcesInfo == nil {
		response.BadRequestWithMessage(w, "no resources")
		return
	}
	if err := auth.ZipPostValidator.Validate(resourcesInfo); err != nil {
		response.Unauthorized(w, err)
		return
	}

	srcFiles, resourcesNotFound := files.CollectSrcFileInfos(resourcesInfo)
	if len(resourcesNotFound) > 0 {
		log.Info("Src files NOT found:", resourcesNotFound)
		response.URLsNotFound(w, resourcesNotFound)
		return
	}
	log.Info("Src files found: ", srcFiles)

	if zipKey, err := saveZipInfo(srcFiles); err != nil {
		msg := "Failed to save zip info"
		log.Error(msg, err.Error())
		response.ServerError(w, msg)
	} else {
		log.Info("Zip key: ", zipKey)
		response.ZipKeyAsResponse(w, zipKey)
	}
}

func extractInfoFrom(r *http.Request) *files.ResourcesInfo {
	decoder := json.NewDecoder(r.Body)
	info := files.ResourcesInfo{}
	err := decoder.Decode(&info)
	if err != nil {
		log.Error("failed to decode json", err.Error())
	}

	return &info
}

type ZipInfo struct {
	FileInfos []files.FileInfo `json:"file_infos"`
}

func saveZipInfo(files []files.FileInfo) (key string, err error) {
	info := ZipInfo{
		FileInfos:files,
	}
	marshaled, err1 := json.Marshal(info)
	if err1 != nil {
		err = err1
	} else {
		key = utils.Md5(marshaled)
		err = utils.GetRedisClient().SaveKey(key, marshaled, StatusDuration)
	}
	return
}