package main

import (
	"fmt"
	"os"
	"sinaugo/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	app.Use(logger.New(), cors.New())
	routes.Routes(app)
	
	envLoader()

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