package main

import (
	"log"
	"net/http"
)

const (
	PORT = ":8080"
)

func main() {

	routes := routes()
	log.Println("Starting webserver on ", PORT)
	_ = http.ListenAndServe(PORT, routes)
}
