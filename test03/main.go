package main

import (
	"test03/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/beef/summary", handler.GetQuantityMeatHandler)
	app.Listen(":8000")
}
