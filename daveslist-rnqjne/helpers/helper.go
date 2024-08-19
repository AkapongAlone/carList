package helpers

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateUUIDV4() string {
	uuid, _ := uuid.NewRandom()
	return uuid.String()
}

func CheckIsImage(fileHeader *multipart.FileHeader) (bool, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Read the first 512 bytes to determine the content type
	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return false, err
	}

	// Use http.DetectContentType to determine the content type
	contentType := http.DetectContentType(buf)

	// Check if the content type starts with "image/"
	if contentType != "" && strings.HasPrefix(contentType, "image/") {
		return true, nil
	}

	return false, nil
}

func GetUserID(c *gin.Context) uint {
	userID, exists := c.MustGet("userID").(float64)
	if !exists {
		return 0 // Return 0 if user ID does not exist in the context
	}
	return uint(userID)
}

func GetRoleID(c *gin.Context) uint {
	// fmt.Println(c.Keys)
	roleID, ifExist := c.MustGet("roleID").(float64)
	if !ifExist {
		return 0
	}
	return uint(roleID)
}

func GetIDParam(c *gin.Context) (uint, error) {
	IDStr := c.Param("id")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		return 0, err
	} else if ID < 1 {
		return 0, errors.New("invalid param id ")
	}
	return uint(ID), nil
}
