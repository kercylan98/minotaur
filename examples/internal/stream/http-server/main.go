package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		_, err := ctx.WriteString("pong")
		return err
	})
}
