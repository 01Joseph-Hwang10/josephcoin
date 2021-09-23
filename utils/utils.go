package utils

import (
	"fmt"
	"log"
	"net/http"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func WithLog(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		handler(rw, r)
		logRequest(r)
	}
}

func logRequest(r *http.Request) {
	fmt.Printf("[%s] \"%s\"\n", r.Method, r.RequestURI)
}
