package main

import (
	"anibal/server/middlewares/serve"
)

func main() {
	app := serve.Anibal()

	app.Anibal()
	// Prints: http://localhost:4200
	app.Anibal(serve.ToListen(80))
	// Prints: http://localhost:80
}
