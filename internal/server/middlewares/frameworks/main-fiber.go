package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	route := fiber.New()

	route.Get("/v1/api/properties", getProperties)
	route.Get("/v1/api/properties/{id}", getProperty)
	route.Post("/v1/api/properties", createProperty)
	route.Put("/v1/api/properties/{id}", updateProperty)
	route.Delete("/v1/api/properties/{id}", deleteProperty)

	log.Fatal(route.Listen(":8880"))
}
