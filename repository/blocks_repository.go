package repository

import "todo-go/model"

func BlocksGetAll() []model.BLock {
	user1 := model.BLock{Id: 1, Text: "Serg"}
	user2 := model.BLock{Id: 2, Text: "Pasha"}
	users := []model.BLock{user1, user2}
	return users
}

func BlocksGetById(id *int64) model.BLock {
	return model.BLock{Id: *id, Text: "name"}
}
