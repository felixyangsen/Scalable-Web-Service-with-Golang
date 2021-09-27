package main

import (
	"myapp/database"
	"myapp/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
