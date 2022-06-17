package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to Blog API!")
	godotenv.Load(".env")
	apiPortRaw := os.Getenv("API_PORT")
	apiPort, err := strconv.Atoi(apiPortRaw)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err = app.Listen(fmt.Sprintf(":%d", apiPort))
	if err != nil {
		log.Fatal(err)
	}
}
