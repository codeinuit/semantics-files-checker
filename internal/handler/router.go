package handler

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/codeinuit/semantics-files-checker/internal/models"
	"github.com/codeinuit/semantics-files-checker/internal/usecase"
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

	var resp models.UploadResultResponse
	resp.Students = usecase.CheckZipFilesSemantics(h.Log, zr.File)
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
