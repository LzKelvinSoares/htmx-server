package handler

import (
	"htmx-server/controller"
	"htmx-server/shared/constants"
	"net/http"
)

func GetBooksRoutes(mux *http.ServeMux) {

	// region GET
	mux.HandleFunc(constants.GetRoute(constants.Books), controller.GetBooks)
	mux.HandleFunc(constants.GetRoute(constants.BookById), controller.GetBookById)
	// endregion

	// region POST
	mux.HandleFunc(constants.PostRoute(constants.BooksSearch), controller.GetFilteredBooks)
	mux.HandleFunc(constants.PostRoute(constants.Books), controller.PostBooks)
	// endregion

	// region PUT
	mux.HandleFunc(constants.PutRoute(constants.BookById), controller.PostBookById)
	// endregion

	// region DELETE
	mux.HandleFunc(constants.DeleteRoute(constants.BookById), controller.DeleteBook)
	// endregion
}
