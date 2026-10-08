package handler

import (
	"htmx-server/controller"
	"htmx-server/shared/constants"
	"net/http"
)

func GetBooksRoutes(mux *http.ServeMux) {

	mux.HandleFunc(constants.PostRoute(constants.BooksSearch), controller.GetFilteredBooks)
	mux.HandleFunc(constants.GetRoute(constants.Books), controller.GetBooks)
	mux.HandleFunc(constants.GetRoute(constants.BookById), controller.GetBookById)
	mux.HandleFunc(constants.PostRoute(constants.Books), controller.PostBooks)
	mux.HandleFunc(constants.PutRoute(constants.BookById), controller.PostBookById)
	mux.HandleFunc(constants.DeleteRoute(constants.BookById), controller.DeleteBook)
}
