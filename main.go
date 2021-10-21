package main

import (
	"github.com/01Joseph-Hwang10/josephcoin/cli"
	"github.com/01Joseph-Hwang10/josephcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
