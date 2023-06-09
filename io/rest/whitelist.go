package rest

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
)

func createWhitelist(c echo.Context) error {
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

	w := new(Whitelist)
	if err := c.Bind(w); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(w); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	id, err := manager.CreateWhitelist(c.Request().Context(), w, phone)
	if err != nil {
		if errors.Is(err, manager.ErrDuplicateEntity) || errors.Is(err, manager.ErrNoAccess) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	w.FID = id
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "whitelist created",
	})
}

func getWhitelists(c echo.Context) error {
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

	Whitelists, err := manager.GetWhitelists(c.Request().Context(), phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"whitelists": toWhitelistResSlice(Whitelists)})
}

func deleteWhitelist(c echo.Context) error {
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

	whiteListID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "",
		})
	}

	err = manager.DeleteWhitelist(c.Request().Context(), whiteListID, phone)
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
		"message": "Whitelist deleted successfully",
	})
}
