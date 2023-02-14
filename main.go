package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/to_do_crud/todos/routes"
)

func main() {
	r := mux.NewRouter()
	routes.SetUpRoutes(r)

	fmt.Println("Starting the server on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
