package controller

import (
	"html/template"
	"htmx-server/service"
	"htmx-server/shared/components"
	"htmx-server/shared/types"
	"htmx-server/shared/utils"
	"htmx-server/views"
	"net/http"
	"strconv"
	"strings"
)

var tpl = template.Must(
	template.Must(
		template.ParseFS(views.FS, "books/*.html"),
	).ParseFS(components.FS, "*.html"),
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := service.GetBooks()
	utils.GetHtmlHeader(w)
	tpl.ExecuteTemplate(w, "booklist", books)
}

func GetFilteredBooks(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.FormValue("search"))
	books := service.FindBookByTitle(search)

	utils.GetHtmlHeader(w)
	tpl.ExecuteTemplate(w, "booklist", books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusBadRequest,
			Err:       err,
			CustomMsg: "id inválido",
		}
		utils.GetError(errorParams)
		return
	}

	book := service.FindBookById(id)
	utils.GetHtmlHeader(w)
	if book.Title == "" {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusBadRequest,
			CustomMsg: "Título não encontrado",
		}
		utils.GetError(errorParams)
	}

	data := types.FormData{ID: id, Title: book.Title, Author: book.Author, Editing: true}
	if err := tpl.ExecuteTemplate(w, "edit", data); err != nil {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusInternalServerError,
			Err:       err,
		}
		utils.GetError(errorParams)
	}
}

func PostBooks(w http.ResponseWriter, r *http.Request) {
	book := service.AddBook(w, r)

	utils.GetHtmlHeader(w)
	if book.Title == "" {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusInternalServerError,
		}
		utils.GetError(errorParams)
	}

	if err := tpl.ExecuteTemplate(w, "item", book); err != nil {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusInternalServerError,
			Err:       err,
		}
		utils.GetError(errorParams)
	}
}

func PostBookById(w http.ResponseWriter, r *http.Request) {
	book := service.AddBookById(w, r)

	utils.GetHtmlHeader(w)
	if book.Title == "" {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusInternalServerError,
		}
		utils.GetError(errorParams)
	}

	if err := tpl.ExecuteTemplate(w, "item", book); err != nil {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusInternalServerError,
			Err:       err,
		}
		utils.GetError(errorParams)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusBadRequest,
			Err:       err,
			CustomMsg: "id inválido",
		}
		utils.GetError(errorParams)
		return
	}

	if !service.DeleteBook(id) {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusNotFound,
			CustomMsg: "livro não encontrado",
		}
		utils.GetError(errorParams)
		return
	}

	w.WriteHeader(http.StatusOK) // corpo vazio: o HTMX remove o <li>
}
