package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type handler struct {}

func (h *handler) getHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

func main() {
	e := echo.New()
	
	h := &handler{}

	e.GET("/health", h.getHealth)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}