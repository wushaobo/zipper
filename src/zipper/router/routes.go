package router

import (
	"net/http"
	"zipper/controller/handler"
)

type Route struct {
	Path   string
	Method string
	Action func(http.ResponseWriter, *http.Request)
}

func routes() []Route {
	return []Route {
		Route{
			Path:   "/zip/{key}",
			Method: http.MethodGet,
			Action: handler.DownloadZip,
		},
		Route{
			Path:   "/zip-info",
			Method: http.MethodPost,
			Action: handler.CreateZipInfo,
		},
	}
}