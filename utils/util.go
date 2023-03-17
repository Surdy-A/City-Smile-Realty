package utils

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadImages(c *gin.Context) ([]string, error) {
	var uploadImages []string

	form, err := c.MultipartForm()

	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return nil, err
	}

	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, "./assets/images/uploads/"+filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return nil, err
		}

		uploadImages = append(uploadImages, filename)

	}
	return uploadImages, nil
}
