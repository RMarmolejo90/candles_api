package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/candles_api/internal/database"
)

func main() {
	// attempt database connection
	database.ConnectDB()

	// start server
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is Running!!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
