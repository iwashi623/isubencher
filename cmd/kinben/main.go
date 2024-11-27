package main

import (
	"fmt"
	"os"

	"github.com/iwashi623/kinben"
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

	app, err := kinben.NewKinben(
		port,
		isuconName,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := app.StartServer(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
