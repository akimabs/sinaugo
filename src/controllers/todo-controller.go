package controller

import (
	"sinaugo/src/services"
	"sinaugo/src/utils/flag"

	"github.com/gofiber/fiber/v2"
)

// TodoController -> the propose of todo controller
type TodoController struct {
	Service services.TodoService
}


// TController -> user controller instance
func TController() TodoController {
	return TodoController{
		Service: services.TService(),
	}
}

// GetTodos -> get users routes
// GET /users
func (t *TodoController) GetTodos(res *fiber.Ctx) error{
	todos := t.Service.GetTodos()
	
	return	res.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data": todos,
		"message": flag.GetTodoSuccess.Message,
	})
}