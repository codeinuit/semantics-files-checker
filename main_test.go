package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/codeinuit/semantics-files-checker/internal/handler"
	"github.com/codeinuit/semantics-files-checker/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	r := handler.NewRouter(&handler.Handler{Log: logrus.New()})

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpload(t *testing.T) {
	r := handler.NewRouter(&handler.Handler{Log: logrus.New()})

	buf := bytes.NewBuffer(nil)
	bodyWriter := multipart.NewWriter(buf)

	fw, err := bodyWriter.CreateFormFile("file", "export.zip")
	require.NoError(t, err)

	f, err := os.Open("./testdata/export.zip")
	require.NoError(t, err)

	defer f.Close()

	_, err = io.Copy(fw, f)
	require.NoError(t, err)

	contentType := bodyWriter.FormDataContentType()

	bodyWriter.Close()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/upload", buf)
	require.NoError(t, err)
	req.Header.Set("Content-Type", contentType)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var res models.UploadResultResponse
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&res))
	req.Body.Close()

	assert.Len(t, res.Students, 3)
}
