package rest

import (
	"errors"
	"github.com/google/uuid"
	"net/http"

	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
)

func carEnter(c echo.Context) error {
	l := new(Log)
	if err := c.Bind(l); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	parkingUUIDStr := c.Param("uuid")
	parkingUUID, err := uuid.Parse(parkingUUIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	id, err := manager.CarEnter(c.Request().Context(), l, parkingUUID)
	if err != nil {
		if errors.Is(err, manager.ErrDuplicateEntity) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		if errors.Is(err, manager.ErrInvalidCarTag) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, toLogRes(l, id))
}

func carExit(c echo.Context) error {
	ce := new(Log)
	err := c.Bind(ce)
	if err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	parkingUUIDStr := c.Param("uuid")
	parkingUUID, err := uuid.Parse(parkingUUIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.CarExit(c.Request().Context(), parkingUUID, ce.CarTag())
	if err != nil {
		lg.Error(err)
		if errors.Is(err, manager.ErrNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Car exited successfully",
	})
}
