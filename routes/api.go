package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/todos/todos"
)

func SetUpRoutes(r *mux.Router) {
	r.HandleFunc("/todos", todos.Index).Methods("GET")
	r.HandleFunc("/todos", todos.Create).Methods("POST")
	r.HandleFunc("/todos/{id}", todos.Show).Methods("GET")
	r.HandleFunc("/todos/{id}", todos.Update).Methods("PUT")
	r.HandleFunc("/todos/{id}", todos.Delete).Methods("DELETE")
}
