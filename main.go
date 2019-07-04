package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := GetPort()
	log.Println("[-] Listening on...", port)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(port) // listen and serve on port
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
