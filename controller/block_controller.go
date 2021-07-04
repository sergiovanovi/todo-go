package controller

import (
	"net/http"
	"todo-go/service"
)

func BlocksGetAll(w http.ResponseWriter, r *http.Request) {
	AbstractGetAll(w, service.BlocksGetAll)
}

func BlocksGetById(w http.ResponseWriter, r *http.Request) {
	AbstractGetById(w, r, service.BlocksGetById)
}
