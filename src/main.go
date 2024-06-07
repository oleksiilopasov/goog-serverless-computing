package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"onefor.fun/gosmarty/handlers"
)

func main() {
	router := gin.Default()

	// Route for file upload
	router.POST("/upload", handlers.UploadHandler)

	// Route for liveness probe
	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Start the server
	router.Run(":8080")
}
