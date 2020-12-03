package main

import (
	"fmt"
	"os"
	db "sinaugo/src/database"
	"sinaugo/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(500).SendString(err.Error())
			},
		},
	)

	envLoader()
	db.Connection()
	
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))
	routes.Routes(app)

	app.Listen(":"+ os.Getenv("HOST_PORT"))
}

func envLoader () error {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return err
}