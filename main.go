package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/developers", developers)
	app.Listen(":9000")
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}
