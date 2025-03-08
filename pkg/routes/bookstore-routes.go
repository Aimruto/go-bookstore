package routes

import (
	"github.com/Aimruto/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var ResgisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("DELETE")
}
