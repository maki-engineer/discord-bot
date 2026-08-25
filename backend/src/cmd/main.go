package main

import (
	"discord-bot/src/server/route"
)

func main() {
	r := route.SetupRoutes()

	r.Run(":8080")
}
