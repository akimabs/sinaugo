package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Routes buat manggil semua enpoin disini
func Routes(app *fiber.App){

	app.Use(logger.New(), cors.New())

	app.Get("/", func(res *fiber.Ctx) error {
		return res.SendString("API Sinaugo")
	  })

}