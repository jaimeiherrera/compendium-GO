package main

import (
	db "compendium-go/src/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var a = db.SQLite{}
		print(a)
		c.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":9001")
}
