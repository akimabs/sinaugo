package utils

import (
	"github.com/gofiber/fiber/v2"
)

// ResponseSuccess : returning json structur for success request
func ResponseSuccess(res *fiber.Ctx, data interface{}, flag string) error {
	return res.JSON(fiber.Map{
		"status":  200,
		"message": flag,
		"data":    data,
	})
}

// ResponseNotFound : returning json structur for notfound request
func ResponseNotFound(res *fiber.Ctx, data interface{}, flag string) error {
	return res.JSON(fiber.Map{
		"status":  404,
		"message": flag,
	})
}

// ResponseError : returning json structur for error request
func ResponseError(res *fiber.Ctx, data interface{}, flag string) error {
	return res.JSON(fiber.Map{
		"status":  500,
		"message": flag,
		"data":    data,
	})
}