package routes

import (
	controller "sinaugo/src/controllers"

	"github.com/gofiber/fiber/v2"
)

// Routes buat manggil semua enpoin dimari
func Routes(app *fiber.App){
	controller := controller.TController()


	app.Get("/", func(res *fiber.Ctx) error {
		return res.JSON(fiber.Map{
			"message": "🐣 v1",
		})
	  })

	  route := app.Group("/api/v1")
	  route.Get("/", controller.GetTodos)
	//   route.Post("/", controllers.Store)
	//   route.Get("/:id", controllers.Show)
	//   route.Put("/:id", controllers.Patch)
	//   route.Delete("/:id", controllers.Destroy)
}