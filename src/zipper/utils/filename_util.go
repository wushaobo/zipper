package utils

import "regexp"

const (
	FILE_NAME_PATTERN = "^(.+)(\\.[^\\.\\/]+)$"
)

func FileNameAndExt(filename string) (name string, ext string) {
	filenameRegexp := regexp.MustCompile(FILE_NAME_PATTERN)
	if filenameRegexp.MatchString(filename) {
		subMatches := filenameRegexp.FindStringSubmatch(filename)
		name = subMatches[1]
		ext = subMatches[2]
	} else {
		name = filename
		ext = ""
	}
	return
}
