package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/01Joseph-Hwang10/josephcoin/explorer"
	"github.com/01Joseph-Hwang10/josephcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to josephcoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		fmt.Println("Starting in REST mode...")
		rest.Start(*port)
	case "html":
		fmt.Println("Starting in explorer mode...")
		explorer.Start(*port)
	default:
		usage()
	}
}
