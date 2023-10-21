package main

import (
	"fmt"
	"net/http"
)

func MessageServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "He did a great job today. Attaboy!")
}
