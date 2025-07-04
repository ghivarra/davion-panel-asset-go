package upload

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

type FormUpload struct {
	FolderPath string          `form="folder_path" binding="required"`
	File       *multipart.File `form="file" binding="required"`
}

func Index(c *gin.Context) {
	var form FormUpload
	errForm := c.ShouldBind(form)
	if errForm != nil {
		if os.Getenv("ENV") == "development" {
			c.AbortWithStatusJSON(400, gin.H{
				"status":  "error",
				"message": fmt.Sprintf("%v", errForm),
			})
		} else {
			c.AbortWithStatusJSON(400, gin.H{
				"status":  "error",
				"message": "Failed to upload file/files. Reason: form error.",
			})
		}
	}
}
