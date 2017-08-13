package auth

import (
	"zipper/files"
	"sort"
	"strings"
	"strconv"
)

const (
	TOKEN_EXPIRES_DURATION_MS = 7200*1000
)

var (
	ZipPostValidator = &zipPostValidator{}
)

type zipPostValidator struct {}

func (this *zipPostValidator) Validate(resourcesInfo *files.ResourcesInfo) (err error) {
	if tokenErr := this.validateToken(resourcesInfo); tokenErr != nil {
		err = tokenErr
	} else if timestampErr := this.validateTimestamp(resourcesInfo); timestampErr != nil {
		err = timestampErr
	}

	return
}

func (this *zipPostValidator) validateToken(resourcesInfo *files.ResourcesInfo) (err error) {
	seed := this.extractSeedForToken(resourcesInfo)
	return validateToken(seed, resourcesInfo.Token)
}

func (this *zipPostValidator) extractSeedForToken(resourcesInfo *files.ResourcesInfo) string {
	urls := this.getUrlsFrom(resourcesInfo.FileInfos)
	sort.Strings(urls)
	return strings.Join(urls, "") + resourcesInfo.Timestamp
}

func (this *zipPostValidator) validateTimestamp(resourcesInfo *files.ResourcesInfo) (err error) {
	timestamp, _ := strconv.ParseInt(resourcesInfo.Timestamp, 10, 64)
	return validateTimestamp(timestamp)
}

func (this *zipPostValidator) getUrlsFrom(fileInfos []files.FileInfo) []string {
	urls := []string{}
	for _, info := range fileInfos {
		urls = append(urls, info.URL)
	}
	return urls
}