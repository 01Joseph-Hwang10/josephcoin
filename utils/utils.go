package utils

import (
	"bytes"
	"encoding/gob"
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

func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
}
