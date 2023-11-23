package main

import "github.com/gofiber/fiber/v2"

// Routes
func homeRoute(c *fiber.Ctx) error {
	return c.Render("base", fiber.Map{})
}
