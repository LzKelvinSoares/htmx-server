package main

import (
	"htmx-server/handler"
	"net/http"
	"sync"
)

func main() {
	mux := http.NewServeMux()

	var wg sync.WaitGroup
	wg.Add(1)
	go handler.GetContent(mux, &wg)
	wg.Wait()

	handler.GetBooksRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
