package rest

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
)

func createZone(c echo.Context) error {
	z := new(Zone)
	if err := c.Bind(z); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(z); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	id, err := manager.CreateZone(c.Request().Context(), z)
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
	return c.JSON(http.StatusCreated, toZoneRes(z, id))
}

func getZones(c echo.Context) error {

	_, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
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

	AC := new(entity.WhitelistGetReq)
	if err := c.Bind(AC); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	Zones, err := manager.GetZones(c.Request().Context(), AC.AdminCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, toZoneResSlice(Zones))
}

func updateZone(c echo.Context) error {
	z := new(Zone)
	if err := c.Bind(z); err != nil {
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
	z.FId = int(pid)

	err = manager.UpdateZone(c.Request().Context(), z)
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

	return c.JSON(http.StatusCreated, toZoneRes(z, -1))
}

func deleteZone(c echo.Context) error {
	ZoneID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "",
		})
	}

	zD := new(entity.ZoneDelete)
	if err := c.Bind(zD); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.DeleteZone(c.Request().Context(), ZoneID, zD.AdminUUID)
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
		"message": "Zone deleted successfully",
	})
}

func EnterZone(c echo.Context) error {
	zid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.EnterZone(c.Request().Context(), int(zid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Updated successfully"})
}

func ExitZone(c echo.Context) error {
	zid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.ExitZone(c.Request().Context(), int(zid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Updated successfully"})
}
