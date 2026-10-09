package handler

import (
	"html/template"
	"htmx-server/shared"
	"htmx-server/shared/components"
	"htmx-server/shared/constants"
	"htmx-server/shared/types"
	"htmx-server/views"
	"log"
	"net/http"
	"sync"
)

var tpl = template.Must(
	template.Must(
		template.ParseFS(views.FS, "*.html", "books/*.html"),
	).ParseFS(components.FS, "*.html"),
)

func GetContent(mux *http.ServeMux, wg *sync.WaitGroup) {
	defer wg.Done()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		data := struct{ Form types.FormData }{} // Editing=false: modo criar
		w.Header().Set("Content-Type", constants.CONTENT_TYPE["html"])
		if err := tpl.ExecuteTemplate(w, "index", data); err != nil {
			log.Println("erro no template:", err)
		}
	})

	mux.HandleFunc("GET /styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, shared.FS, "styles/index.css")
	})

}
