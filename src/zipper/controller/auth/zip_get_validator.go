package auth

import (
	"strconv"
	"fmt"
	"zipper/controller/model"
)

var (
	ZipGetValidator = &zipGetValidator{}
)

type zipGetValidator struct {}

func (this *zipGetValidator) Validate(queryInfo *model.QueryInfo) (err error) {
	if tokenErr := this.validateToken(queryInfo); tokenErr != nil {
		err = tokenErr
	} else if timestampErr := this.validateTimestamp(queryInfo); timestampErr != nil {
		err = timestampErr
	}
	return
}

func (this *zipGetValidator) validateToken(queryInfo *model.QueryInfo) (err error) {
	seed := this.extractSeedForToken(queryInfo)
	return validateToken(seed, queryInfo.Token)
}

func (this *zipGetValidator) extractSeedForToken(queryInfo *model.QueryInfo) string {
	return fmt.Sprintf("%s?timestamp=%s", queryInfo.URI, queryInfo.Timestamp)
}

func (this *zipGetValidator) validateTimestamp(queryInfo *model.QueryInfo) (err error) {
	timestamp, _ := strconv.ParseInt(queryInfo.Timestamp, 10, 64)
	return validateTimestamp(timestamp)
}
