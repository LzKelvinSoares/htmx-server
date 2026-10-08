package constants

import "strings"

const (
	// RouteHome is the route for the home page
	RouteHome   = "/"
	Books       = "/books"
	BooksSearch = "/books/search"
	BookById    = "/books/{id}"
)

var ROUTES = map[string]string{
	RouteHome:   RouteHome,
	Books:       Books,
	BooksSearch: BooksSearch,
	BookById:    BookById,
}

func GetRoute(routeName string) string {
	return strings.Join([]string{"GET ", ROUTES[routeName]}, " ")
}

func PostRoute(routeName string) string {
	return strings.Join([]string{"POST ", ROUTES[routeName]}, " ")
}

func PutRoute(routeName string) string {
	return strings.Join([]string{"PUT ", ROUTES[routeName]}, " ")
}

func DeleteRoute(routeName string) string {
	return strings.Join([]string{"DELETE ", ROUTES[routeName]}, " ")
}
