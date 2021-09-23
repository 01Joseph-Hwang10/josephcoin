package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	Blockchain "github.com/01Joseph-Hwang10/josephcoin/blockchain"
	"github.com/01Joseph-Hwang10/josephcoin/utils"
	"github.com/gorilla/mux"
)

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(apiSpec))
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := Blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == Blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
		return
	}
	encoder.Encode(block)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(Blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		Blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}
