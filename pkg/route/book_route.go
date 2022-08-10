package route

import (
	"github.com/FarrukhMahkamov/book_management/pkg/controller"
	"github.com/gorilla/mux"
)

// Routes for Books
var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books", controller.GetAllBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controller.DeleteBook).Methods("DELETE")
}
