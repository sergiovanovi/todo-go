package service

import (
	"todo-go/model"
	"todo-go/repository"
)

func UsersGetAll() ([]interface{}, error) {
	users := repository.UsersGetAll()
	return usersConvert(users), nil
}

func UsersGetById(id *int64) (interface{}, error) {
	return repository.UsersGetById(id), nil
}

func UsersAdd(user interface{}) (interface{}, error) {
	//u := user.(model.User)
	//u.Id = 1
	//user = u
	return user, nil
}

func usersConvert(users []model.User) []interface{} {
	var arr []interface{}
	for _, u := range users {
		arr = append(arr, u)
	}
	return arr
}
