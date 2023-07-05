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

// swagger:route PUT /v1/whitelist/approve/{id} Parking_Admin approveWhitelist
//
// # This route is used to approve a whitelist by ID
//
// responses:
//
//	200: ErrorMessage
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func approveWhitelist(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	c.Response().Header().Set("session_token", sessionToken)
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

	err = manager.ApproveWhitelist(c.Request().Context(), whiteListID, phone)
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
		"message": "Whitelist approved successfully",
	})
}

// swagger:route POST /v1/user/whitelist/request User accessRequest
//
// # This route is used by user to request access for a parking
//
// responses:
//
//	201: ErrorMessage
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func requestWhitelist(c echo.Context) error {
	phone, sessionToken, err := authenticateUser(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	c.Response().Header().Set("session_token", sessionToken)
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
		if errors.Is(err, manager.ErrDuplicateEntity) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		if errors.Is(err, manager.ErrParkingNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
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

// swagger:route GET /v1/whitelists/approved Parking_Admin getApprovedWhitelists
//
// # This route is used by parking admin to get approved whitelists
//
// responses:
//
//	200: ApprovedWhitelistsGetRes
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func getWhitelistsApproved(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	c.Response().Header().Set("session_token", sessionToken)
	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	Whitelists, err := manager.GetWhitelists(c.Request().Context(), phone, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"whitelists": toWhitelistOfficeResSlice(Whitelists)})
}

// swagger:route GET /v1/whitelists/toApprove Parking_Admin getWhitelistsToApprove
//
// # This route is used by parking admin to get access requests
//
// responses:
//
//	200: WhitelistsToApproveGetRes
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func getWhitelistsToApprove(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	c.Response().Header().Set("session_token", sessionToken)
	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	Whitelists, err := manager.GetWhitelists(c.Request().Context(), phone, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"whitelists": toWhitelistOfficeResSlice(Whitelists)})
}

// swagger:route DELETE /v1/whitelist/{id} Parking_Admin deleteWhitelist
//
// # This route is used to delete a whitelist by ID
//
// responses:
//
//	200: ErrorMessage
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func deleteWhitelist(c echo.Context) error {
	phone, sessionToken, err := authenticateParkingAdmin(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	c.Response().Header().Set("session_token", sessionToken)
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

// swagger:route GET /v1/user/whitelists/requested User getUserRequests
//
// # This route is used to by user to see their requests
//
// responses:
//
//	200: UserRequestsRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func getUserWhitelists(c echo.Context) error {
	phone, sessionToken, err := authenticateUser(c.Request().Context(), c.Request().Header.Get("session_token"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	c.Response().Header().Set("session_token", sessionToken)
	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	Whitelists, err := manager.GetUserWhitelists(c.Request().Context(), phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"whitelists": toWhitelistUserResSlice(Whitelists)})
}
