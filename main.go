package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo-go/controller"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", controller.UsersGetAll).Methods(http.MethodGet)
	r.HandleFunc("/users", controller.UsersAdd).Methods(http.MethodPost)
	r.HandleFunc("/users/{id:[0-9]+}", controller.UsersGetById).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", controller.UsersGetById).Methods(http.MethodPut)
	r.HandleFunc("/users/{id:[0-9]+}", controller.UsersGetById).Methods(http.MethodDelete)

	r.HandleFunc("/blocks", controller.BlocksGetAll).Methods(http.MethodGet)
	r.HandleFunc("/blocks/{id:[0-9]+}", controller.BlocksGetById).Methods(http.MethodGet)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
