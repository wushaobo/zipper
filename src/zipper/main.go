package main

import (
	"zipper/router"
	"zipper/log"
)

func main() {
	if err := router.HttpServe(); err != nil {
		log.FatalAndExit(err)
	}
}