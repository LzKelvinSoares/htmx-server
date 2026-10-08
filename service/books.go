package service

import (
	"htmx-server/shared/types"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	mu     sync.Mutex
	nextID = 3
	books  = []types.Book{}
)

func GetBooks() []types.Book {
	mu.Lock()
	defer mu.Unlock()
	dataMap := GetDataMap()
	booksData := []types.Book{
		{ID: "1", Title: "The Final Empire", Author: "Brandon Sanderson"},
		{ID: "2", Title: "The Way of Kings", Author: "Brandon Sanderson"},
	}
	dataMap["booksData"] = booksData
	books = dataMap["booksData"]
	return append([]types.Book(nil), dataMap["booksData"]...)
}

func AddBook(w http.ResponseWriter, r *http.Request) types.Book {
	mu.Lock()
	defer mu.Unlock()
	title := strings.TrimSpace(r.FormValue("title"))
	author := strings.TrimSpace(r.FormValue("author"))

	if title == "" || author == "" {
		http.Error(w, "título e autor são obrigatórios", http.StatusBadRequest)
		return types.Book{}
	}

	book := types.Book{
		ID:     strconv.Itoa(nextID),
		Title:  title,
		Author: author,
	}
	nextID++
	dataMap := GetDataMap()
	dataMap["booksData"] = append(dataMap["booksData"], book)
	return book
}

func AddBookById(w http.ResponseWriter, r *http.Request) types.Book {
	mu.Lock()
	defer mu.Unlock()
	title := strings.TrimSpace(r.FormValue("title"))
	author := strings.TrimSpace(r.FormValue("author"))
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return types.Book{}
	}

	if title == "" || author == "" {
		http.Error(w, "título e autor são obrigatórios", http.StatusBadRequest)
		return types.Book{}
	}

	book := types.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}

	dataMap := GetDataMap()
	for i, b := range dataMap["booksData"] {
		if b.ID == id {
			dataMap["booksData"][i] = book
		}
	}
	return book
}

func FindBookById(id int) types.Book {
	mu.Lock()
	defer mu.Unlock()
	dataMap := GetDataMap()
	for _, b := range dataMap["booksData"] {
		if b.ID == strconv.Itoa(id) {
			return b
		}
	}
	return types.Book{}
}

func FindBookByTitle(search string) []types.Book {
	mu.Lock()
	defer mu.Unlock()
	dataMap := GetDataMap()
	var filteredBooks []types.Book
	for _, b := range dataMap["booksData"] {
		if strings.Contains(b.Title, search) {
			filteredBooks = append(filteredBooks, b)
		}
	}
	return filteredBooks
}

func DeleteBook(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	dataMap := GetDataMap()
	for i, b := range dataMap["booksData"] {
		if b.ID == strconv.Itoa(id) {
			dataMap["booksData"] = append(dataMap["booksData"][:i], dataMap["booksData"][i+1:]...)
			return true
		}
	}
	return false
}
