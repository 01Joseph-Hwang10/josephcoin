package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/01Joseph-Hwang10/josephcoin/utils"
	"github.com/gorilla/mux"
)

func attachHandlers(handler *mux.Router) {
	// Iterating
	for i := 0; i < len(urlData); i++ {
		// Collect Methods
		var methods []string
		for _, v := range urlData[i].Methods {
			methods = append(methods, v.Method)
		}
		// Attach
		handler.HandleFunc(urlData[i].urlPattern, utils.WithLog(urlData[i].handler)).Methods(methods...)
	}
}

func generateApiSpec() {
	data, err := json.Marshal(urlData)
	apiSpec = data
	utils.HandleErr(err)
}

func bootstrap() {
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	attachHandlers(router)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
