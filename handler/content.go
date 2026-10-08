package handler

import (
	"htmx-server/shared"
	"htmx-server/views"
	"net/http"
	"sync"
)

func GetContent(mux *http.ServeMux, wg *sync.WaitGroup) {
	defer wg.Done()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, views.FS, "index.html")
	})

	mux.HandleFunc("GET /styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, shared.FS, "styles/index.css")
	})

}
