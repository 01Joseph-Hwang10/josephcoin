package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"
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
	if !Production {
		currentTime := time.Now()
		fmt.Printf("%s [%s] \"%s\"\n", currentTime, r.Method, r.RequestURI)
	}
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

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func DevLog(i interface{}) {
	if !Production {
		currentTime := time.Now()
		fmt.Printf("%s [LOG] %s\n\n", currentTime, i)
	}
}
