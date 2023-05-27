package rest

import (
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
	"net/http"
	"time"
)

func parkingAdminSignIn(c echo.Context) error {
	cr := new(entity.Credentials)
	if err := c.Bind(cr); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	sessionToken, err := manager.GetParkingAdminPasswordByPhone(c.Request().Context(), *cr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
	return c.NoContent(200)
}
