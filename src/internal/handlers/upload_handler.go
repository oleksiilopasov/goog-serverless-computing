package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"onefor.fun/gosmarty/internal/database"
	"onefor.fun/gosmarty/internal/storage"
	"onefor.fun/gosmarty/pkg/config"
)

type UploadHandler struct {
	cfg *config.Config
}

func NewUploadHandler(cfg *config.Config) gin.HandlerFunc {
	handler := &UploadHandler{cfg: cfg}
	return handler.Upload
}

func (h *UploadHandler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	defer file.Close()

	// Upload the file to Cloud Storage
	objectName := header.Filename
	if err := storage.UploadToCloudStorage(h.cfg.CloudStorageBucket, objectName, file); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error uploading file: %s", err.Error()))
		return
	}

	// Record information about the uploaded file in PostgreSQL
	if err := database.RecordUploadInDatabase(h.cfg, objectName, header.Size); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error recording upload in database: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "File uploaded successfully!")
}
