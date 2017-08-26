package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var DB = make(map[string]string)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world!")
	})

	r.Run(":" + port)
}
