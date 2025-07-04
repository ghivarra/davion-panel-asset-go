package imageController

import (
	"fmt"
	stdimage "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/ghivarra/davion-panel-asset-go/common"
	"github.com/ghivarra/davion-panel-asset-go/environment"
	"github.com/ghivarra/davion-panel-asset-go/module/library"
	"github.com/gin-gonic/gin"
)

type Query struct {
	Width    string `form:"width"`
	Height   string `form:"height"`
	Priority string `form:"priority"`
}

func decidePriority(options Query) string {
	widthLen := len(options.Width)
	heightLen := len(options.Height)
	priorityLen := len(options.Priority)

	if priorityLen > 0 {
		return options.Priority
	}

	if widthLen == 0 && heightLen > 0 {
		return "height"
	}

	if heightLen == 0 && widthLen > 0 {
		return "width"
	}

	if heightLen == 0 && widthLen == 0 && priorityLen == 0 {
		return "real"
	}

	// default is width
	return "width"
}

func buildDistFileName(options Query, filePath string) string {
	i := strings.LastIndex(filePath, ".")
	exts := filePath[i:]
	name := filePath[:i]

	extendedName := "_"

	if len(options.Width) > 1 {
		extendedName += fmt.Sprintf("w-%s_", options.Width)
	}

	if len(options.Height) > 1 {
		extendedName += fmt.Sprintf("h-%s_", options.Height)
	}

	extendedName += fmt.Sprintf("p-%s_", options.Priority)

	// return
	return name + extendedName + exts
}

func encodeImage(output io.Writer, image stdimage.Image, format imaging.Format) error {
	formatStr := format.String()
	var errEncoding error
	switch formatStr {
	case "JPEG":
		errEncoding = jpeg.Encode(output, image, &jpeg.Options{Quality: 90})
	case "JPG":
		errEncoding = jpeg.Encode(output, image, &jpeg.Options{Quality: 90})
	case "PNG":
		errEncoding = png.Encode(output, image)
	case "GIF":
		errEncoding = gif.Encode(output, image, &gif.Options{})
	}

	if errEncoding != nil {
		if environment.ENV == "development" {
			fmt.Println(errEncoding)
		}
		return errEncoding
	}

	return nil
}

func Get(c *gin.Context) {
	var options Query
	errQuery := c.ShouldBindQuery(&options)
	if errQuery != nil {
		library.SendErrorResponse(c, 400, "Failed to get image. Reason: Bad Query")
		return
	}

	// decide priority
	options.Priority = decidePriority(options)

	// if priority is forced then width and height should be supplied
	if options.Priority == "forced" {
		if len(options.Width) < 1 || len(options.Height) < 1 {
			library.SendErrorResponse(c, 400, "Failed to process image. Reason: If the priority is forced then the width and height options should be supplied.")
			return
		}
	}

	// get file
	reqPath := c.Param("path")
	reqPath = strings.ReplaceAll(reqPath, "\\", "/")
	filePath := fmt.Sprintf("%s/upload/image/%s", common.ROOTPATH, reqPath)

	// send error if file not exist
	if !library.FileExist(filePath) {
		if environment.ENV == "development" {
			fmt.Println(filePath)
		}
		library.SendErrorResponse(c, 400, "File not exist.")
		return
	}

	// send file if priority is real
	if options.Priority == "real" {
		c.File(filePath)
		return
	}

	// check if width and height is valid
	if len(options.Width) > 0 && !slices.Contains(environment.ALLOWED_IMAGE_WIDTH, options.Width) {
		library.SendErrorResponse(c, 400, "Failed to process image. Reason: Width is not allowed.")
		return
	}

	if len(options.Height) > 0 && !slices.Contains(environment.ALLOWED_IMAGE_HEIGHT, options.Height) {
		library.SendErrorResponse(c, 400, "Failed to process image. Reason: Height is not allowed.")
		return
	}

	// form dist name and folder
	distFolder := fmt.Sprintf("%s/upload/dist/image", common.ROOTPATH)
	distName := buildDistFileName(options, reqPath)

	if strings.Contains(distName[1:], "/") {
		i := strings.LastIndex(distName[1:], "/")
		subFolder := distName[0 : i+1]

		// modify folder and name
		distFolder = fmt.Sprintf("%s%s", distFolder, subFolder)
		distName = distName[i+1:]
	}

	// set dist path
	distPath := fmt.Sprintf("%s%s", distFolder, distName)

	// check if file already exist
	// return the already existed file
	if library.FileExist(distPath) {
		c.File(distPath)
		return
	}

	// set options as int
	width, _ := strconv.Atoi(options.Width)
	height, _ := strconv.Atoi(options.Height)

	switch options.Priority {
	case "width":
		height = 0
	case "height":
		width = 0
	}

	// load image
	image, errOpeningImage := imaging.Open(filePath, imaging.AutoOrientation(true))
	if errOpeningImage != nil {
		if environment.ENV == "development" {
			fmt.Println(errOpeningImage)
		}
		library.SendErrorResponse(c, 400, "Failed to process image. Reason: File is not an image or file is corrupt.")
		return
	}

	// create all folders if not exist
	os.MkdirAll(distFolder, 0755)

	// write new file
	output, errWriting := os.Create(distPath)
	if errWriting != nil {
		if environment.ENV == "development" {
			fmt.Println(errWriting)
		}
		library.SendErrorResponse(c, 400, "Failed to process image. Reason: Cannot write new file.")
		return
	}
	defer output.Close()

	// start resizing images
	var imageResized *stdimage.NRGBA

	if options.Priority != "forced" && width > 0 && height > 0 {

		imageResized = imaging.Fit(image, width, height, imaging.NearestNeighbor)

	} else {

		imageResized = imaging.Resize(image, width, height, imaging.NearestNeighbor)
	}

	// get original image format
	imageFormat, _ := imaging.FormatFromFilename(reqPath)

	// start encoding
	errEncoding := encodeImage(output, imageResized, imageFormat)
	if errEncoding != nil {
		if environment.ENV == "development" {
			fmt.Println(errEncoding)
		}
		library.SendErrorResponse(c, 400, "Failed to process image. Reason: Cannot encoding the file.")
		return
	}

	// send output
	c.File(distPath)
}
