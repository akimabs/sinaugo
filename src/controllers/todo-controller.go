package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Index return all Todos
func Index(res *fiber.Ctx) error {
	return res.JSON(&fiber.Map{
		"status": true,
		"books":  "asd",
	})
}