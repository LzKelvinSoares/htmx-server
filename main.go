package main

import (
	_ "embed"
	"htmx-server/handler"
	"log"
	"net/http"
	"os"
	"sync"
)

var indexHTML []byte

func main() {
	mux := http.NewServeMux()

	var wg sync.WaitGroup
	wg.Add(1)
	go handler.GetContent(mux, &wg)
	wg.Wait()

	handler.GetBooksRoutes(mux)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
