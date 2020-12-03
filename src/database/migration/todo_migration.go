package migration

import (
	"sinaugo/src/database/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// CreateTodo is create user tabel for migration
func CreateTodo(conn *gorm.DB){	
	conn.AutoMigrate(&entity.Todo{})
	logrus.Info("Success running migration")
}