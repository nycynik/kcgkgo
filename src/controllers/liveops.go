package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

// Structs
type LiveOpsStatus struct {
	Status  string `json:"status" xml:"status"`
	DB string `json:"db" xml:"db"`
}

// Methods
func LiveOpsSetup(e *echo.Echo) {
	e.GET("/liveops/ping", handlePing)
	e.GET("/liveops/status", handleStatus)
}

func handlePing(c echo.Context) error {
	return c.String(http.StatusOK, "pong\n")
}

func handleStatus(c echo.Context) error {
	s := &LiveOpsStatus{
		Status:  "Okay",
		DB: "Okay",
	}
	return c.JSON(http.StatusOK, s)
}
