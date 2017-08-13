package utils

import (
	"encoding/hex"
	"crypto/md5"
)

func Md5(content []byte) string {
	sum := md5.Sum(content)
	return hex.EncodeToString(sum[:])
}

