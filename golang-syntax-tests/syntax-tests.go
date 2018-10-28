package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/learning-go/golang-syntax-tests/controllers"
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

	fmt.Printf("Server started at: %s:%d\n", host, port)
	controllers.LoadRoutes()
	log.Fatal(server.ListenAndServe())
}

func executeMiddleware() {
	fmt.Printf("Middleware called!\n")
}
