package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
)

type Handler struct {
	cgiHandler *cgi.Handler
}

func NewHandler(reposDir, backendCommand string) *Handler {
	return &Handler{
		&cgi.Handler{
			Path: "/bin/sh",
			Args: []string{"-c", backendCommand},
			Dir:  ".",
			Env: []string{
				fmt.Sprintf("GIT_PROJECT_ROOT=%v", reposDir),
				"GIT_HTTP_EXPORT_ALL=1",
			},
		},
	}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println(req)
	h.cgiHandler.ServeHTTP(rw, req)
}

func main() {
	var (
		reposDir       = flag.String("r", ".", "Directory containing git repositories")
		backendCommand = flag.String("c", "git http-backend", "CGI binary to execute")
		addr           = flag.String("l", ":80", "Address/port to listen on")
	)
	flag.Parse()

	http.Handle("/", NewHandler(*reposDir, *backendCommand))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
