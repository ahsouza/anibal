package main

import (
	ani "anibal/internal/server/middlewares/http"
)

func main() {
	//ani.Default()
	// Print: http://localhost:4200
	//ani.Default(ani.SetPort(3226))
	// Print: http://localhost:3226

	ani.NewServer(ani.ToListen(":8891"), ani.Framework("fiber"))

}
