package controller

import (
	"net/http"
	"reflect"
	"todo-go/model"
	"todo-go/service"
)

func BlocksGetAll(w http.ResponseWriter, r *http.Request)  {
	AbstractGetAll(w, service.BlocksGetAll)
}

func BlocksGetById(w http.ResponseWriter, r *http.Request)  {
	AbstractGetById(w, r, service.BlocksGetById)
}

func BlocksAdd(w http.ResponseWriter, r *http.Request)  {
	AbstractAdd(w, r, reflect.TypeOf(model.BLock{}), service.BlocksAdd)
}
