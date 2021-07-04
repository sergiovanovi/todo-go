package service

import (
	"todo-go/model"
	"todo-go/repository"
)

func BlocksGetAll() ([]interface{}, error) {
	blocks := repository.BlocksGetAll()
	return blocksConvert(blocks), nil
}

func BlocksGetById(id *int64) (interface{}, error) {
	return repository.BlocksGetById(id), nil
}

func BlocksAdd(block interface{}) (interface{}, error) {
	//b := block.(model.Block)
	//b.Id = 1
	//block = u
	return block, nil
}

func blocksConvert(users []model.BLock) []interface{} {
	var arr []interface{}
	for _, u := range users {
		arr = append(arr, u)
	}
	return arr
}