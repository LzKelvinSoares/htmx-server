package controller

import (
	"html/template"
	"htmx-server/service"
	"htmx-server/shared/constants"
	"htmx-server/views/books"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tpl = template.Must(template.ParseFS(books.FS, "*.html"))

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := service.GetBooks()
	w.Header().Set("Content-Type", constants.CONTENT_TYPE["html"])
	tpl.ExecuteTemplate(w, "booklist", books)
}

func GetFilteredBooks(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.FormValue("search"))
	books := service.FindBookByTitle(search)

	w.Header().Set("Content-Type", constants.CONTENT_TYPE["html"])
	tpl.ExecuteTemplate(w, "booklist", books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	book := service.FindBookById(id)
	w.Header().Set("Content-Type", constants.CONTENT_TYPE["html"])
	if book.Title == "" {
		http.Error(w, "erro interno", http.StatusInternalServerError)
	}

	if err := tpl.ExecuteTemplate(w, "edit", book); err != nil {
		log.Println("erro no template:", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
	}
}

func PostBooks(w http.ResponseWriter, r *http.Request) {
	book := service.AddBook(w, r)

	w.Header().Set("Content-Type", constants.CONTENT_TYPE["html"])
	if book.Title == "" {
		http.Error(w, "erro interno", http.StatusInternalServerError)
	}

	if err := tpl.ExecuteTemplate(w, "item", book); err != nil {
		log.Println("erro no template:", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
	}
}

func PostBookById(w http.ResponseWriter, r *http.Request) {
	book := service.AddBookById(w, r)

	w.Header().Set("Content-Type", constants.CONTENT_TYPE["html"])
	if book.Title == "" {
		http.Error(w, "erro interno", http.StatusInternalServerError)
	}

	if err := tpl.ExecuteTemplate(w, "item", book); err != nil {
		log.Println("erro no template:", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	if !service.DeleteBook(id) {
		http.Error(w, "livro não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK) // corpo vazio: o HTMX remove o <li>
}
