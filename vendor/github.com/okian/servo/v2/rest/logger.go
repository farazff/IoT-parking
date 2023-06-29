package rest

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
)

func logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		start := time.Now()
		if err = next(c); err != nil {
			lg.Errorf("%s %s %d %s Error: %q", c.Request().Method,
				c.Request().URL.String(),
				c.Response().Status,
				time.Since(start).String(),
				err.Error())
			c.Error(err)
			return
		}
		lg.Infof("%s %s %d %s", c.Request().Method, c.Request().URL.String(), c.Response().Status, time.Since(start))

		return
	}
}
