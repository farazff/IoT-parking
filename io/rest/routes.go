package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Running")
}
