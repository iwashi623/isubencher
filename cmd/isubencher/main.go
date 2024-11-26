package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/iwashi623/isubencher"
)

func main() {
	port := os.Getenv("PORT")
	s := http.Server{
		Addr: ":" + port,
	}

	isuconName, err := isubencher.GetIsuconName()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/bench", isubencher.BenchHandler)

	fmt.Println("ISUCON_NAME: " + isuconName)
	fmt.Println("Server is running on port " + port)
	s.ListenAndServe()

}
