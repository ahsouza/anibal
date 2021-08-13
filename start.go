package main

import (
	ani "anibal/internal/server/middlewares/http"
)

func main() {
	ani.Server()
	// Prints: http://localhost:4200
	ani.Server(ani.ToListen(80))
	// Prints: http://localhost:80
}
