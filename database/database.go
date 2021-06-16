package database

import (
	"github.com/mstzn/modanisa_backend/config"
	"github.com/mstzn/modanisa_backend/models"
)

type Database interface {
	GetAll() []models.ToDo
	GetOneById(id string) models.ToDo
	Insert(do models.ToDo)bool
}

type DbConfig struct {

}

func GetDatabase() Database {
	driver := config.GetEnvironmentVariable("DATABASE_DRIVER", "memory")
	if driver == "memory" {
		return Memory{}
	}

	return Memory{}
}