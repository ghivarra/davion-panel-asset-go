package environment

import (
	"os"
	"strings"
)

// main env
var ENV string

// server env
var SERVER_HOST string
var SERVER_PORT string

// app env
var TEMPORARY_FOLDER string
var UPLOAD_FOLDER string
var ALLOWED_HOST []string
var ALLOWED_FILE_MIME []string

// img env
var ALLOWED_IMAGE_WIDTH []string
var ALLOWED_IMAGE_HEIGHT []string

func Save() {

	// main env
	ENV = os.Getenv("ENV")

	// server env
	SERVER_HOST = os.Getenv("SERVER_HOST")
	SERVER_PORT = os.Getenv("SERVER_PORT")

	// app env
	TEMPORARY_FOLDER = os.Getenv("TEMPORARY_FOLDER")
	UPLOAD_FOLDER = os.Getenv("UPLOAD_FOLDER")
	ALLOWED_HOST = strings.Split(os.Getenv("ALLOWED_HOST"), "|")
	ALLOWED_FILE_MIME = strings.Split(os.Getenv("ALLOWED_FILE_MIME"), "|")

	// image env
	ALLOWED_IMAGE_WIDTH = strings.Split(os.Getenv("ALLOWED_IMAGE_WIDTH"), "|")
	ALLOWED_IMAGE_HEIGHT = strings.Split(os.Getenv("ALLOWED_IMAGE_HEIGHT"), "|")
}
