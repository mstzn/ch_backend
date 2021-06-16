package database

import "github.com/mstzn/modanisa_backend/models"

var ToDos[]models.ToDo

type Memory struct {
	DbConfig
}

func (d Memory) GetAll() []models.ToDo {
	return ToDos
}

func (d Memory) GetOneById(id string) models.ToDo {
	return ToDos[0]
}

func (d Memory) Insert(do models.ToDo)bool {
	ToDos = append([]models.ToDo{do}, ToDos...)

	return true
}
