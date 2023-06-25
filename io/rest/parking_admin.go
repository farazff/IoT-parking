package rest

import (
	"errors"
	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
	"net/http"
	"strconv"
	"time"
)

func createParkingAdmin(c echo.Context) error {
	_, sessionToken, err := authenticateSystemAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	p := new(ParkingAdmin)
	if err := c.Bind(p); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	id, err := manager.CreateParkingAdmin(c.Request().Context(), p)
	if err != nil {
		if errors.Is(err, manager.ErrDuplicateEntity) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, toParkingAdminRes(p, id))
}

func getParkingAdmin(c echo.Context) error {
	ParkingAdminID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error in ParkingAdmin id",
		})
	}

	ParkingAdmin, err := manager.GetParkingAdmin(c.Request().Context(), ParkingAdminID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, toParkingAdminRes(ParkingAdmin, -1))
}

func getParkingAdmins(c echo.Context) error {
	_, sessionToken, err := authenticateSystemAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	ParkingAdmins, err := manager.GetParkingAdmins(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, toParkingAdminResSlice(ParkingAdmins))
}

func updateParkingAdmin(c echo.Context) error {
	_, sessionToken, err := authenticateSystemAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	p := new(ParkingAdmin)
	if err := c.Bind(p); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	pid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	p.FID = int(pid)

	err = manager.UpdateParkingAdmin(c.Request().Context(), p)
	if err != nil {
		if errors.Is(err, manager.ErrDuplicateEntity) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		if errors.Is(err, manager.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, toParkingAdminRes(p, -1))
}

func deleteParkingAdmin(c echo.Context) error {
	_, sessionToken, err := authenticateSystemAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	ParkingAdminID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "",
		})
	}
	err = manager.DeleteParkingAdmin(c.Request().Context(), ParkingAdminID)
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
		"message": "ParkingAdmin deleted successfully",
	})
}
