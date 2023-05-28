package rest

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"

	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
)

func createZone(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
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

	id, parkingID, err := manager.CreateZone(c.Request().Context(), z, phone)
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
	z.FParkingID = parkingID
	return c.JSON(http.StatusCreated, echo.Map{"zone": toZoneRes(z, id)})
}

func getZones(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
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

	zones, err := manager.GetZones(c.Request().Context(), phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"zones": toZoneResSlice(zones)})
}

func getZone(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
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

	zoneID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	zone, err := manager.GetZone(c.Request().Context(), int(zoneID), phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"zone": toZoneRes(zone, -1)})
}

func updateZone(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
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

	z := new(Zone)
	if err := c.Bind(z); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	zoneID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	z.FID = int(zoneID)

	err = manager.UpdateZone(c.Request().Context(), z, phone)
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

	return c.JSON(http.StatusCreated, echo.Map{"zone": toZoneRes(z, -1)})
}

func deleteZone(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
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

	ZoneID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "",
		})
	}

	err = manager.DeleteZone(c.Request().Context(), ZoneID, phone)
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

	parkingUUID := c.Param("uuid")
	_, err = uuid.Parse(parkingUUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.EnterZone(c.Request().Context(), int(zid), parkingUUID)
	if err != nil {
		if errors.Is(err, manager.ErrNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
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

	parkingUUID := c.Param("uuid")
	_, err = uuid.Parse(parkingUUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.ExitZone(c.Request().Context(), int(zid), parkingUUID)
	if err != nil {
		if errors.Is(err, manager.ErrNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Updated successfully"})
}
