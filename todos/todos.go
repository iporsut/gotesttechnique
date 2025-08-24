package main

import (
	"net/http"
)

func main() {
	server := NewServer()
	http.ListenAndServe(":8080", server)
}
