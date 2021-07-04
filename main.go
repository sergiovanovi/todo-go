package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo-go/controller"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/blocks", controller.BlocksGetAll).Methods(http.MethodGet)
	r.HandleFunc("/blocks/{id:[0-9]+}", controller.BlocksGetById).Methods(http.MethodGet)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
