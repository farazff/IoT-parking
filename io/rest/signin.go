package rest

import (
	_ "embed"
	"errors"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/manager"
	"github.com/go-openapi/runtime/middleware"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
	"github.com/okian/servo/v2/rest"
	"net/http"
	"time"
)

//go:embed swagger.yaml
var swagger []byte

func docs() error {
	ops := middleware.RedocOpts{
		Path:    "/swagger",
		SpecURL: "/swagger.yaml",
	}
	rest.EchoGet("/swagger", echo.WrapHandler(middleware.Redoc(ops, nil)))
	rest.EchoGet("/swagger.yaml", func(c echo.Context) error {
		c.Response().Write(swagger)
		return nil
	})
	return nil
}

// swagger:route POST /v1/systemAdmin/signIn System_Admin systemAdminSignIn
//
// This route is used by system admin to sign in.
//
// responses:
//
//	204: NoContent
//	401: ErrorMessage
//	500: ErrorMessage
func systemAdminSignIn(c echo.Context) error {
	cr := new(entity.Credentials)
	if err := c.Bind(cr); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(cr); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	sessionToken, err := manager.GetSystemAdminPasswordByPhone(c.Request().Context(), *cr)
	if err != nil {
		if errors.Is(err, manager.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
	return c.NoContent(http.StatusNoContent)
}

// swagger:route POST /v1/parkingAdmin/signIn Parking_Admin parkingAdminSingIn
//
// This route is used by parking admin to sign in.
//
// responses:
//
//	204: NoContent
//	401: ErrorMessage
//	500: ErrorMessage
func parkingAdminSignIn(c echo.Context) error {
	cr := new(entity.Credentials)
	if err := c.Bind(cr); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(cr); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	sessionToken, err := manager.GetParkingAdminPasswordByPhone(c.Request().Context(), *cr)
	if err != nil {
		if errors.Is(err, manager.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
	return c.NoContent(http.StatusNoContent)
}

// swagger:route POST /v1/user/signIn User userSingIn
//
// This route is used by user to sign in.
//
// responses:
//
//	204: NoContent
//	401: ErrorMessage
//	500: ErrorMessage
func userSignIn(c echo.Context) error {
	cr := new(entity.Credentials)
	if err := c.Bind(cr); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(cr); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	sessionToken, err := manager.GetUserPasswordByPhone(c.Request().Context(), *cr)
	if err != nil {
		if errors.Is(err, manager.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
	return c.NoContent(http.StatusNoContent)
}

// swagger:route POST /v1/user/signUp User userSingUp
//
// This route is used by user to sign up.
//
// responses:
//
//	204: NoContent
//	404: ErrorMessage
//	500: ErrorMessage
func userSignUp(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(user); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	err := manager.CreateUser(c.Request().Context(), user)
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
	return c.NoContent(http.StatusNoContent)
}
