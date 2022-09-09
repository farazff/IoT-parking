package rest

import (
	"errors"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
	"net/http"
)

func createWhitelist(c echo.Context) error {
	p := new(WhitelistCreateReq)
	if err := c.Bind(p); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	id, err := manager.CreateWhitelist(c.Request().Context(), p, p.AdminCode)
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
	return c.JSON(http.StatusCreated, toWhitelistRes(p, id))
}

func getWhitelists(c echo.Context) error {
	wgr := new(entity.WhitelistGetReq)
	err := c.Bind(wgr)
	if err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	Whitelists, err := manager.GetWhitelists(c.Request().Context(), *wgr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, toWhitelistResSlice(Whitelists))
}

func deleteWhitelist(c echo.Context) error {
	wdr := new(entity.WhitelistDeleteReq)
	err := c.Bind(wdr)
	if err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.DeleteWhitelist(c.Request().Context(), *wdr)
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
