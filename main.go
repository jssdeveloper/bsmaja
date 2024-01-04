package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("base", fiber.Map{})
	})

	// log.Fatal(app.Listen("192.168.0.164:3100")) // dev mac mini
	log.Fatal(app.Listen("192.168.0.180:3100")) // dev macbook
	// log.Fatal(app.Listen(":80")) // prod
}
