package rest

import (
	"errors"
	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
	"net/http"
	"strconv"
)

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Running")
}

func getParking(c echo.Context) error {
	parkingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error in parking id",
			"status":  http.StatusBadRequest,
		})
	}

	parking, err := manager.GetParking(c.Request().Context(), parkingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}
	return c.JSON(http.StatusOK, toParkingRes(parking))
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

func deleteParking(c echo.Context) error {
	parkingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "",
			"status":  http.StatusBadRequest,
		})
	}
	err = manager.DeleteParking(c.Request().Context(), parkingID)
	if err != nil {
		lg.Error(err)
		if errors.Is(err, manager.ErrParkingNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Parking deleted successfully",
	})
}
