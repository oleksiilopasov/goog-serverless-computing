package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"onefor.fun/gosmarty/internal/handlers"
	"onefor.fun/gosmarty/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	router := gin.Default()

	// Route for info page
	router.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a Go web application for testing purposes.")
	})

	// Route for file upload
	router.POST("/upload", handlers.NewUploadHandler(cfg))

	// Route for liveness probe
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Start the server
	router.Run(":8080")
}
