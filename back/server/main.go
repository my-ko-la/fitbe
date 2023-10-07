package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("../../front/views", ".html")

    app := fiber.New(fiber.Config{
		Views: engine,
	})

    app.Get("/", func (c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
			"Title": "Fitbe Studio",
		})
    })

    app.Get("/bookings", func (c *fiber.Ctx) error {
        return c.Render("booking-form", fiber.Map{
			"Title": "Fitbe Studio",
		})
    })

    log.Fatal(app.Listen(":6969"))
}