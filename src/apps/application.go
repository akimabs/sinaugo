package apps

import (
	db "sinaugo/src/database"
	"sinaugo/src/database/migration"

	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Application -> application instance
type Application struct {
}

// CreateApp -> method to create gin application
func (a Application) CreateApp(res *fiber.App){
	res.Use(middleware.Logger(), cors.New())
	configureAppDB()
}

/**
 * configuration database application
 */
func configureAppDB() {
	db.AppConnection()
}

/**
 * configuration database application for testing
 */
 func configureTestDB() {
	db.TestConnection()
	conn := db.GetDB()
	migration.CreateTodo(conn)
}