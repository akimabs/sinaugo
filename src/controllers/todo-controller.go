package controllers

import (
	"sinaugo/src/utils/flag"
	utils "sinaugo/src/utils/response"

	"github.com/gofiber/fiber/v2"
)

// GetData will be show all data
type GetData struct {
	ID int `json:"id"`
	Title string `json:"title"`
	IsDone bool `json:"isDone"`
}

// Index return all Todos
func Index(res *fiber.Ctx) error {

	data := [] GetData{
		{
			ID: 1,
			Title: "Belajar masak",
			IsDone: false,
		},
	}
	return utils.ResponseSuccess(res, data, flag.GetAllTodo.Message )
}