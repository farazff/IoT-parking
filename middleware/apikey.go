package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func ApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, key := range strings.Split(viper.GetString("api-key"), ",") {
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

func AdminApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, key := range strings.Split(viper.GetString("admin_api-key"), ",") {
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
