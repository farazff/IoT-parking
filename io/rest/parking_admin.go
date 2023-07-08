package rest

import (
	"errors"
	"github.com/farazff/IoT-parking/manager"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
	"net/http"
	"strconv"
)

// swagger:route POST /v1/parkingAdmin System_Admin createParkingAdmin
//
// # This route is used to create parking admin
//
// responses:
//
//	201: ParkingAdminCreateRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func createParkingAdmin(c echo.Context) error {
	p := new(ParkingAdmin)
	if err := c.Bind(p); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(p); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	id, err := manager.CreateParkingAdmin(c.Request().Context(), p)
	if err != nil {
		if errors.Is(err, manager.ErrDuplicateEntity) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		if errors.Is(err, manager.ErrParkingNotFound) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{"parking_admin": toParkingAdminRes(p, id)})
}

// swagger:route GET /v1/parkingAdmin/{id} System_Admin getParkingAdmin
//
// # This route is used to get a single parking admin by ID
//
// responses:
//
//		200: ParkingAdminGetRes
//		400: ErrorMessage
//		401: ErrorUnauthorizedMessage
//	 404: ErrorMessage
//		500: ErrorMessage
func getParkingAdmin(c echo.Context) error {
	ParkingAdminID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error in ParkingAdmin id",
		})
	}

	ParkingAdmin, err := manager.GetParkingAdmin(c.Request().Context(), ParkingAdminID)
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
	return c.JSON(http.StatusOK, echo.Map{"parking_admin": toParkingAdminRes(ParkingAdmin, -1)})
}

// swagger:route GET /v1/parkingAdmins System_Admin getParkingAdmins
//
// # This route is used to get all parking admins
//
// responses:
//
//	200: ParkingAdminsGetRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func getParkingAdmins(c echo.Context) error {
	ParkingAdmins, err := manager.GetParkingAdmins(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"parking_admins": toParkingAdminResSlice(ParkingAdmins)})
}

// swagger:route PUT /v1/parkingAdmin/{id} System_Admin updateParkingAdmin
//
// # This route is used to update a parking admin
//
// responses:
//
//	201: ParkingAdminUpdateRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func updateParkingAdmin(c echo.Context) error {
	p := new(ParkingAdmin)
	if err := c.Bind(p); err != nil {
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
	p.FID = int(pid)

	err = manager.UpdateParkingAdmin(c.Request().Context(), p)
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
		if errors.Is(err, manager.ErrParkingNotFound) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{"parking_admin": toParkingAdminRes(p, -1)})
}

// swagger:route DELETE /v1/parkingAdmin/{id} System_Admin deleteParkingAdmin
//
// # This route is used to delete a parking admin by ID
//
// responses:
//
//	200: ErrorMessage
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func deleteParkingAdmin(c echo.Context) error {
	ParkingAdminID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "",
		})
	}
	err = manager.DeleteParkingAdmin(c.Request().Context(), ParkingAdminID)
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
		"message": "ParkingAdmin deleted successfully",
	})
}
