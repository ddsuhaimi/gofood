package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// Retrieve the hostname of the server
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		// Retrieve server identifier from environment variable
		serverID := os.Getenv("SERVER_ID")
		if serverID == "" {
			serverID = "unknown-server"
		}

		c.JSON(200, gin.H{
			"message":  "Hello, World!",
			"hostname": hostname,
			"serverID": serverID,
		})
	})
	router.Run(":8080")
}
