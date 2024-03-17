package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
