package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// Retrieve the hostname of the server
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		c.JSON(200, gin.H{
			"message":  "Hello, World!",
			"hostname": hostname,
		})
	})
	router.Run(":8080")
}
