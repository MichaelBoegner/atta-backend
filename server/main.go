package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(MessageServer)
	log.Fatal(http.ListenAndServe(":4141", handler))
}
