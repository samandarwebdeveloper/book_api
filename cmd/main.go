package main

import (
	"books/pkg/api/books"
	"books/pkg/api/categories"
	"books/pkg/config"
	"books/pkg/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	books.RegisterRoutes(app, h)
	categories.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
