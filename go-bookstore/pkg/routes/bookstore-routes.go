package routes

import (
	"github.com/anishgowda21/go-projects/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.Createbook).Methods("Post")
	router.HandleFunc("/book/",controllers.GetBook).Methods("Get")
	router.HandleFunc("/book/{bookId}",controllers.GetBookByID).Methods("Get")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}