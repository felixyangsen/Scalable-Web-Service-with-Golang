package main

import (
	"myapp/database"
	"myapp/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":" + os.Getenv("PORT"))
}
