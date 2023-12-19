package main

import (
	"goworktoday/database"
	"goworktoday/urls"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.InitDB()
	urls.RegisterRoutes(r, db)
	r.Run(":8080")
}
