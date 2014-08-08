package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
)

func main() {
	var (
		reposDir       = flag.String("r", ".", "Directory containing git repositories")
		backendCommand = flag.String("c", "git http-backend", "CGI binary to execute")
		addr           = flag.String("l", ":80", "Address/port to listen on")
	)
	flag.Parse()

	handler := &cgi.Handler{
		Path: "/bin/sh",
		Args: []string{"-c", *backendCommand},
		Dir:  ".",
		Env: []string{
			fmt.Sprintf("GIT_PROJECT_ROOT=%v", *reposDir),
			"GIT_HTTP_EXPORT_ALL=1",
		},
	}

	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
