package main

import (
	"fmt"
	"os"

	"github.com/iwashi623/isubencher"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	isuconName := os.Getenv("ISUCON_NAME")
	if isuconName == "" {
		fmt.Println("ISUCON_NAME is not set")
		os.Exit(1)
	}

	app := isubencher.NewIsubencher(
		port,
		isuconName,
	)
	if err := app.StartServer(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
