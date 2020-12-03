package seed

import (
	"sinaugo/src/database/entity"

	"github.com/bxcodec/faker/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// CreateTodo is seeder to input data
func CreateTodo(db *gorm.DB) error {
	todo := entity.Todo{}
	if err := faker.FakeData(&todo); err != nil {
		logrus.Errorln("Error todo seed", err)
	}

	return db.Create(&entity.Todo{
		Name: todo.Name,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdateAt: todo.UpdateAt,
	}).Error
}