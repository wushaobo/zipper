package auth

import (
	"errors"
	"time"
	"zipper/config"
	"zipper/utils"
)

func validateTimestamp(timestamp int64) (err error) {
	currentTimestamp := time.Now().UTC().Unix()
	if timestamp + TOKEN_EXPIRES_DURATION_MS < currentTimestamp {
		err = errors.New("The request has been expired")
	}
	return
}

func validateToken(seed string, givenToken string) (err error) {
	token := buildToken(seed)
	if token != givenToken {
		err = errors.New("The token is invalid in request")
	}
	return
}

func buildToken(seed string) string {
	secret_key := config.Http.SecureKey
	content := []byte(seed + secret_key)

	return utils.Md5(content)[:20]
}
