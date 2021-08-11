package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	route := echo.New()
	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	route.GET("/v1/api/properties", getProperties)
	route.GET("/v1/api/properties/{id}", getProperty)
	route.POST("/v1/api/properties", createProperty)
	route.PUT("/v1/api/properties/{id}", updateProperty)
	route.DELETE("/v1/api/properties/{id}", deleteProperty)

	log.Fatal(route.Start(":8880"))
}
