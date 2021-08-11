package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	v1 := route.Group("/v1")
	{
		v1.GET("/api/properties", getProperties)
		v1.GET("/api/properties/{id}", getProperty)
		v1.POST("/api/properties", createProperty)
		v1.PUT("/api/properties/{id}", updateProperty)
		v1.DELETE("/api/properties/{id}", deleteProperty)
	}

	log.Fatal(route.Run(":8880"))

}
