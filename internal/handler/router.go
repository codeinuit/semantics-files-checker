package handler

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
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

		r, err := file.Open()
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, nil)
		}
		defer r.Close()

		reader, err := ioutil.ReadAll(r)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, nil)
		}

		zr, err := zip.NewReader(bytes.NewReader(reader), file.Size)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, nil)
		}

		students := make(map[string][]string)

		for _, file := range zr.File {
			if !file.FileInfo().IsDir() {
				folder := filepath.Base(filepath.Dir(file.Name))
				fmt.Printf("file %s in folder %s\n", filepath.Base(file.Name), folder)
				students[folder] = append(students[folder], filepath.Base(file.Name))

				continue
			}

			fmt.Printf("folder found: %s\n", file.Name)
		}

		for n, val := range students {
			fmt.Printf("student found: %s with files %s\n", n, val)
		}

		c.JSON(http.StatusOK, nil)
	})

	return r
}
