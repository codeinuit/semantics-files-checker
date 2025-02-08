package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	r := initRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
