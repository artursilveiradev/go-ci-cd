package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/health")
	h := &handler{}

	if assert.NoError(t, h.getHealth(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"Status":"OK"}`, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
