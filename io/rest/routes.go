package rest

import (
	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"net/http"
)

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Running")
}

func getParkings(c echo.Context) error {
	parkings, err := manager.GetParkings(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}
	return c.JSON(http.StatusOK, toParkingResSlice(parkings))
}
