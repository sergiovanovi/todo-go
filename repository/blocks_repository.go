package repository

import (
	"fmt"
	"todo-go/model"
)

func BlocksGetAll() []model.BlockModel {
	db := connect()
	defer db.Close()

	var blocks []model.BlockModel
	err := db.Model(&blocks).Select()
	if err != nil {
		fmt.Println(err)
		return []model.BlockModel{}
	}
	return blocks
}

func BlocksGetById(id *int64) model.BlockModel {
	db := connect()
	defer db.Close()

	block := model.BlockModel{Id: *id}
	err := db.Model(&block).WherePK().Select()
	if err != nil {
		fmt.Println(err)
		return model.BlockModel{}
	}
	return block
}
