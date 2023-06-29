package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func SystemAdminApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, key := range strings.Split(viper.GetString("system_admin_api_key"), ",") {
			if c.Request().Header.Get("api-key") == key {
				if err := next(c); err != nil {
					c.Error(err)
				}
				return nil
			}
		}
		return c.NoContent(http.StatusUnauthorized)
	}
}

func ParkingAdminApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, key := range strings.Split(viper.GetString("parking_admin_api_key"), ",") {
			if c.Request().Header.Get("api-key") == key {
				if err := next(c); err != nil {
					c.Error(err)
				}
				return nil
			}
		}
		return c.NoContent(http.StatusUnauthorized)
	}
}

func HardwareApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, key := range strings.Split(viper.GetString("hardware_api_key"), ",") {
			if c.Request().Header.Get("api-key") == key {
				if err := next(c); err != nil {
					c.Error(err)
				}
				return nil
			}
		}
		return c.NoContent(http.StatusUnauthorized)
	}
}

func UserApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, key := range strings.Split(viper.GetString("user_api_key"), ",") {
			if c.Request().Header.Get("api-key") == key {
				if err := next(c); err != nil {
					c.Error(err)
				}
				return nil
			}
		}
		return c.NoContent(http.StatusUnauthorized)
	}
}
