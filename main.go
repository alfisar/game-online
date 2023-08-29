package main

import "game-online/router"

func main() {
	route := router.NewRouter()
	route.Run(":8080")
}
