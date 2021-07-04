package service

import (
	"todo-go/dto"
	"todo-go/model"
	"todo-go/repository"
)

func BlocksGetAll() ([]interface{}, error) {
	blocks := repository.BlocksGetAll()
	return blocksConvert(arraysModelToDto(blocks)), nil
}

func BlocksGetById(id *int64) (interface{}, error) {
	return modelToDto(repository.BlocksGetById(id)), nil
}

func blocksConvert(users []dto.BLockDto) []interface{} {
	var arr []interface{}
	for _, u := range users {
		arr = append(arr, u)
	}
	return arr
}

func dtoToModel(dto dto.BLockDto) model.BlockModel {
	return model.BlockModel{
		Id:             dto.Id,
		Text:           dto.Text,
		Rate:           dto.Rate,
		CreateDatetime: dto.CreateDatetime,
	}
}

func modelToDto(blockModel model.BlockModel) dto.BLockDto {
	return dto.BLockDto{
		Id:             blockModel.Id,
		Text:           blockModel.Text,
		Rate:           blockModel.Rate,
		CreateDatetime: blockModel.CreateDatetime,
	}
}

func arraysDtoToModel(dtos []dto.BLockDto) []model.BlockModel {
	var models []model.BlockModel
	for _, e := range dtos {
		models = append(models, dtoToModel(e))
	}
	return models
}

func arraysModelToDto(models []model.BlockModel) []dto.BLockDto {
	var dtos []dto.BLockDto
	for _, e := range models {
		dtos = append(dtos, modelToDto(e))
	}
	return dtos
}
