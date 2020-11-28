package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

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