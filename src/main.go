package main

import (
	"github.com/gin-gonic/gin"
	"onefor.fun/gosmarty/handlers"
)

func main() {
	router := gin.Default()

	// Route for file upload
	router.POST("/upload", handlers.UploadHandler)

	// Start the server
	router.Run(":8080")
}
