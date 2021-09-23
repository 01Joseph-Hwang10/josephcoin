package rest

import (
	"fmt"
)

var apiSpec []byte

var port string

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	generateApiSpec()
	bootstrap()
}
