package main

import (
	"flag"

	"github.com/gofiber/fiber/v3"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(map[string]string{"msg": "Server is on fire!"})
	})
	app.Listen(*listenAddr)
}
