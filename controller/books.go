package controller

import (
	"html/template"
	"htmx-server/service"
	"htmx-server/shared/components"
	"htmx-server/shared/constants"
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
	utils.GetTemplate(w, tpl, constants.Booklist, books)
}

func GetFilteredBooks(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.FormValue("search"))
	books := service.FindBookByTitle(search)

	utils.GetTemplate(w, tpl, constants.Booklist, books)
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
	if book.Title == "" {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusBadRequest,
			CustomMsg: "Título não encontrado",
		}
		utils.GetError(errorParams)
	}

	data := types.FormData{ID: id, Title: book.Title, Author: book.Author, Editing: true}
	utils.GetTemplate(w, tpl, constants.Edit, data)
}

func PostBooks(w http.ResponseWriter, r *http.Request) {
	book := service.AddBook(w, r)

	if book.Title == "" {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusInternalServerError,
		}
		utils.GetError(errorParams)
	}

	utils.GetTemplate(w, tpl, constants.Item, book)
}

func PostBookById(w http.ResponseWriter, r *http.Request) {
	book := service.AddBookById(w, r)

	if book.Title == "" {
		errorParams := utils.GetErrorParams{
			W:         w,
			ErrorCode: http.StatusInternalServerError,
		}
		utils.GetError(errorParams)
	}

	utils.GetTemplate(w, tpl, constants.Item, book)
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
