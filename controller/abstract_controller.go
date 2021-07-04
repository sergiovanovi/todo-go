package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"strconv"
)

func AbstractGetAll(w http.ResponseWriter, f func() ([]interface{}, error))  {
	arr, err := f()
	if err == nil {
		er := json.NewEncoder(w).Encode(arr)
		err = er
		if err == nil {
			return
		}
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func AbstractGetById(w http.ResponseWriter, r *http.Request, f func(*int64) (interface{}, error)) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err == nil {
		obj, er := f(&id)
		err = er
		if err == nil {
			err = json.NewEncoder(w).Encode(obj)
			if err == nil {
				return
			}
		}
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func AbstractAdd(w http.ResponseWriter, r *http.Request, t reflect.Type, f func(interface{}) (interface{}, error)) {
	obj := reflect.New(t)
	err := json.NewDecoder(r.Body).Decode(&obj)
	if err == nil {
		fmt.Println(obj)
		u, er := f(obj)
		err = er
		if err == nil {
			err = json.NewEncoder(w).Encode(u)
			if err == nil {
				return
			}
		}
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}