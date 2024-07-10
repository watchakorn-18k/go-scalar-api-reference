package main

import (
	"log"
	"simple_api/middlewares"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	middlewares.Logger(app)
	app.Use("/api/docs", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.yaml",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})

		if err != nil {
			return err
		}
		c.Type("html")
		return c.SendString(htmlContent)
	})
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Hello World",
		})
	})
	log.Fatal(app.Listen(":3000"))
}
