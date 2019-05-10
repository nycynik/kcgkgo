package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func APISetup(e *echo.Echo) {
	e.GET("/api/test", handleCall)
}

func handleCall(c echo.Context) error {
	return c.String(http.StatusOK, "It works!\n")
}

