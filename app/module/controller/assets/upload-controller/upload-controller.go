package uploadController

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/ghivarra/davion-panel-asset-go/common"
	"github.com/ghivarra/davion-panel-asset-go/environment"
	"github.com/ghivarra/davion-panel-asset-go/module/library"
	"github.com/gin-gonic/gin"
)

type FormUpload struct {
	FolderPath string                `form:"path" binding:"required"`
	Name       string                `form:"name" binding:"max=128"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}

func checkMimeType(mime *mimetype.MIME, allowedMime []string) bool {
	result := false
	for _, allowed := range allowedMime {
		if mime.Is(allowed) {
			result = true
			break
		}
	}

	return result
}

func Index(c *gin.Context) {
	var form FormUpload
	errForm := c.ShouldBind(&form)
	if errForm != nil {
		if os.Getenv("ENV") == "development" {
			library.SendErrorResponse(c, 400, fmt.Sprintf("%v", errForm))
		} else {
			library.SendErrorResponse(c, 400, "Failed to upload file/files. Reason: form error.")
		}
		return
	}

	// move uploaded file to dir path
	tempName := library.RandomString(32) + ".TEMP"
	tempPath := fmt.Sprintf("%s/%s/%s", common.ROOTPATH, environment.TEMPORARY_FOLDER, tempName)
	c.SaveUploadedFile(form.File, tempPath)

	// check mime
	detect, errDetection := mimetype.DetectFile(tempPath)
	if errDetection != nil {
		// remove temp file
		os.Remove(tempPath)
		// send response
		if os.Getenv("ENV") == "development" {
			library.SendErrorResponse(c, 400, fmt.Sprintf("%v", errDetection))
		} else {
			library.SendErrorResponse(c, 400, "Failed to upload file/files. Reason: File cannot be recognized.")
		}
		return
	}

	allowedMime := environment.ALLOWED_FILE_MIME

	if !checkMimeType(detect, allowedMime) {
		errMessage := fmt.Sprintf("Failed to upload file/files. Reason: File type is not allowed. Allowed File type: %s", strings.Join(allowedMime, ", "))
		library.SendErrorResponse(c, 400, errMessage)
		return
	}

	// generate random name if name is not supplied
	if len(form.Name) < 1 {
		dotLocation := strings.LastIndex(form.File.Filename, ".")
		extension := form.File.Filename[dotLocation:]
		form.Name = library.RandomString(32) + extension
	}

	// set variables
	var newFilePath string
	var newURIPath string

	// if mime is not image
	// then we send it into public/dist
	// and return
	if strings.HasPrefix(detect.String(), "image") {
		newFilePath = fmt.Sprintf("%s/%s/image/%s/%s", common.ROOTPATH, environment.UPLOAD_FOLDER, form.FolderPath, form.Name)
		newURIPath = fmt.Sprintf("assets/image/%s/%s", form.FolderPath, form.Name)
	} else {
		newFilePath = fmt.Sprintf("%s/public/dist/%s/%s", common.ROOTPATH, form.FolderPath, form.Name)
		newURIPath = fmt.Sprintf("dist/%s/%s", form.FolderPath, form.Name)
	}

	// move uploaded file to new folder based on path
	c.SaveUploadedFile(form.File, newFilePath)

	// remove temp file
	os.Remove(tempPath)

	// return file uri
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "File has been succesfully uploaded.",
		"data": gin.H{
			"uri": newURIPath,
		},
	})
}
