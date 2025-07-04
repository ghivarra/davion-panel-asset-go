package library

import (
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)

	// if not error then file exist
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	// return
	return string(b)
}

func SendErrorResponse(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, gin.H{
		"status":  "error",
		"message": message,
	})
}
