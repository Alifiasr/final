package main

import (
	"MyGram/config"
	"MyGram/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
