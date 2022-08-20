//go:build localserver
// +build localserver

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	dir := os.Getenv("DIR")
	if dir == "" {
		log.Fatalln("DIR must be specified")
	}

	log.Printf("open for the server serving %s: http://localhost:%s\n", dir, port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.FileServer(http.Dir(dir))); err != nil {
		log.Fatalln("Failed to start server", err)
	}
}
