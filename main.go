package main

import (
	"log"
	"net/http"
	"fmt"
)

var port = ":8080"

func main() {
	DictionarySetup()

	router := NewRouter()
	fmt.Println("API Listening")
	log.Fatal(http.ListenAndServe(port, router))
}
