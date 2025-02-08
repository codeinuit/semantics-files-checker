package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		if filepath.Ext(file.Filename) != ".zip" {
			c.JSON(http.StatusUnprocessableEntity, nil)
		}

		c.JSON(http.StatusOK, nil)
	})

	return r
}
