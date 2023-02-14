package routes

import (
	"github.com/gorilla/mux"
	"github.com/to_do_crud/todos/controllers"
)

func SetUpRoutes(r *mux.Router) {
	todosController := &controllers.TodosController{}

	r.HandleFunc("/todos", todosController.Index).Methods("GET")
	r.HandleFunc("/todos", todosController.Create).Methods("POST")
	r.HandleFunc("/todos/{id}", todosController.Show).Methods("GET")
	r.HandleFunc("/todos/{id}", todosController.Update).Methods("PUT")
	r.HandleFunc("/todos/{id}", todosController.Delete).Methods("DELETE")
}
