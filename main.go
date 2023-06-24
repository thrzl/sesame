package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("you forgot the file!")
	}

	filePath := os.Args[1]

	gin.SetMode(gin.ReleaseMode)
	// Initialize Gin-Gonic router
	router := gin.Default()

	// Serve the HTML file
	router.GET("/", func(c *gin.Context) {
		c.File(filePath)
	})

	// Start the server
	port := "8080" // Change this to the desired port number
	address := fmt.Sprintf(":%s", port)
	log.Printf("running server on port %s", address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
