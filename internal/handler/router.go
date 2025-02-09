package handler

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/codeinuit/semantics-files-checker/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Log *logrus.Logger
}

func (h Handler) UploadZip(c *gin.Context) {
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
		slice := strings.Split(file.Name, "/")
		if len(slice) <= 1 {
			continue
		}

		student := slice[1]

		_, ok := students[student]
		if !ok {
			h.Log.Debugf("added path: %s", student)
			students[student] = []string{}
		}

		if !file.FileInfo().IsDir() {
			folder := filepath.Base(filepath.Dir(file.Name))
			h.Log.Infof("file %s in folder %s", filepath.Base(file.Name), folder)
			students[student] = append(students[student], filepath.Base(file.Name))
		}
	}

	var resp models.UploadResultResponse
	for n := range students {
		resp.Students = append(resp.Students, n)
	}

	c.JSON(http.StatusOK, resp)
}

func (h Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func NewRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", h.Ping)
	r.POST("upload", h.UploadZip)

	return r
}
