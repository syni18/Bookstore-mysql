package routes

import (
	"github.com/gorilla/mux"
	"github.com/sahil/bookstore-mysql/pkg/controller"
)

var RegisterBookStoreRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/book/", controller.GetBook).Methods("GET")
	routes.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	routes.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")
	routes.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
}
