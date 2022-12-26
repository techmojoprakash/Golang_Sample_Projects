// https://www.youtube.com/watch?v=hWmR8YtlFlE
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
