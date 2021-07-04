package repository

import "todo-go/model"

func UsersGetAll() []model.User {
	user1 := model.User{Id: 1, Name: "Serg"}
	user2 := model.User{Id: 2, Name: "Pasha"}
	users := []model.User{user1, user2}
	return users
}

func UsersGetById(id *int64) model.User {
	return model.User{Id: *id, Name: "name"}
}
