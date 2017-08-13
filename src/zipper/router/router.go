package router

import (
	"net/http"
	"time"
	"fmt"
	"zipper/config"
	"github.com/gorilla/mux"
	"zipper/controller/response"
	"zipper/log"
)

const (
    READ_TIMEOUT = 30 * time.Second
    WRITE_TIMEOUT = 2 * time.Hour
)

func enableCORS(rw http.ResponseWriter) http.ResponseWriter {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
	    "Accept, Content-Type, Content-Length, Accept-Encoding, X-Requested-With, Authorization")
	return rw
}

func checkMethod(rw http.ResponseWriter, r *http.Request, httpMethod string) bool {
	if r.Method == http.MethodOptions {
		// Stop here if its Preflighted OPTIONS request
		return false
	}

	if r.Method != httpMethod {
		response.MethodNotAllowed(rw)
		return false
	}

	return true
}

func addRoute(router *mux.Router, route Route) {
	decoratedFunc := func(rw http.ResponseWriter, r *http.Request) {
		rw = enableCORS(rw)
		if !checkMethod(rw, r, route.Method) {
			return
		}
		log.Access(r)
		route.Action(rw, r)
	}

	router.Path(route.Path).HandlerFunc(decoratedFunc)
}

func listenAndServe(router *mux.Router) error {
	httpConf := config.Http
	server := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", httpConf.ListenPort),
		ReadTimeout:    READ_TIMEOUT,
		WriteTimeout:   WRITE_TIMEOUT,
		Handler: router,
	}
	return server.ListenAndServe()
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes() {
		addRoute(router, route)
	}

	return router
}

func HttpServe() error {
	return listenAndServe(createRouter())
}