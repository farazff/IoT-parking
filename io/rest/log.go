package rest

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
)

func carEnter(c echo.Context) error {
	carTag := c.Param("tag")
	parkingUUIDStr := c.Param("uuid")
	parkingUUID, err := uuid.Parse(parkingUUIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	_, err = manager.CarEnter(c.Request().Context(), carTag, parkingUUID)
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
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Car exited successfully",
	})
}

func carExit(c echo.Context) error {
	carTag := c.Param("tag")
	parkingUUIDStr := c.Param("uuid")
	parkingUUID, err := uuid.Parse(parkingUUIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.CarExit(c.Request().Context(), parkingUUID, carTag)
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

// swagger:route GET /v1/user/logs/{:page} User getUserLogs
//
// # This route is used by user to get their logs
//
// responses:
//
//	200: UserLogsRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func getUserLogs(c echo.Context) error {
	phone, sessionToken, err := authenticateUser(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	userLogs, err := manager.GetUserLogs(c.Request().Context(), phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"logs": toUserLogsResSlice(userLogs)})
}
