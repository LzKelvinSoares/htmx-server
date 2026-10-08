package handler

import (
	"net/http"
	"sync"
)

func GetContent(mux *http.ServeMux, wg *sync.WaitGroup) {
	defer wg.Done()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	mux.HandleFunc("GET /styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "shared/styles/index.css")
	})

}
