package repository

import (
	db "sinaugo/src/database"
	"time"

	"gorm.io/gorm"
)

// TodoRepository -> the propose of user repository is handling query for user entity
type TodoRepository struct {
	Conn *gorm.DB
}

// URepository -> user repository instance to get user table connection
func URepository() TodoRepository {
	return TodoRepository{Conn: db.GetDB().Table("todos")}
}
// Todos -> define Data
type Todos struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt time.Time `json:"updateAt"`
}

// GetUsers -> method to get all users in database
func (r *TodoRepository) GetUsers() []Todos {
	var todos []Todos
	r.Conn.Select("name, completed, created_at, update_at").Find(&todos)
	return todos
}