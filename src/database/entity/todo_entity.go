package entity

import "time"

// Todo model
type Todo struct {
	ID int `gorm:"primaryKey" json:"id"`
	Name string `gorm:"varchar(50)" json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt time.Time `json:"updateAt"`
};