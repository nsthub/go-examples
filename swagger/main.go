// Sample service with Swagger included.
//
// This is an example of a GoLang based service with Swagger UI included.
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: v1
//     Contact: NST Hub NST Info<info@nsthub.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"go-examples/swagger/routes"
	"os"
	"net/http"
	"fmt"
)

func main() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	server := http.NewServeMux()
	server.HandleFunc("/hello", hello)
	http.Handle("/api", http.StripPrefix("/api", http.FileServer(http.Dir("./swagger-ui"))))
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server))

	routes.UserById()
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintln(w, "Hello, world!")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}