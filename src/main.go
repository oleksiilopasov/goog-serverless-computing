package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"onefor.fun/gosmarty/handlers"
)

func main() {
	router := gin.Default()

	// Route for info page
	router.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a Go web application for testing purposes")
	})

	// Route for file upload
	router.POST("/upload", handlers.UploadHandler)

	// Route for liveness probe
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Start the server
	router.Run(":8080")
}
