package controller

import (
	"net/http"
	"reflect"
	"todo-go/model"
	"todo-go/service"
)

func UsersGetAll(w http.ResponseWriter, r *http.Request)  {
	AbstractGetAll(w, service.UsersGetAll)
}

func UsersGetById(w http.ResponseWriter, r *http.Request)  {
	AbstractGetById(w, r, service.UsersGetById)
}

func UsersAdd(w http.ResponseWriter, r *http.Request)  {
	AbstractAdd(w, r, reflect.TypeOf(model.User{}), service.UsersAdd)
}