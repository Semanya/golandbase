package main

import (
	"app/internal/app/handlers/hello"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/hello", hello.Handler)
	http.ListenAndServe(":9090", nil)
}
