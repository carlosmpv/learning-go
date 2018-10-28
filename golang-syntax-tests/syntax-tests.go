package main

import (
	"fmt"
	"log"
	"net/http"
)

var server http.Server
var host string
var port int

func main() {
	host = "localhost"
	port = 8080

	server = http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: nil,
	}

	fmt.Printf("Server started at: %s:%d", host, port)
	log.Fatal(server.ListenAndServe())
}

func responseHandler(w http.ResponseWriter, r *http.Request) {

}
