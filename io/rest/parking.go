package rest

import (
	"errors"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/manager"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/okian/servo/v2/lg"
	"net/http"
	"strconv"
)

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Running...")
}

// swagger:route POST /v1/parking System_Admin createParking
//
// # This route is used to create parking
//
// responses:
//
//	201: ParkingCreateRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func createParking(c echo.Context) error {
	p := new(Parking)
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

	id, Puuid, err := manager.CreateParking(c.Request().Context(), p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	p.FID = id
	return c.JSON(http.StatusCreated, echo.Map{"parking": toParkingRes(p, 0, Puuid)})
}

// swagger:route GET /v1/parking/{id} System_Admin getParking
//
// # This route is used to get a single parking by ID
//
// responses:
//
//		200: ParkingGetRes
//		400: ErrorMessage
//		401: ErrorUnauthorizedMessage
//	 404: ErrorMessage
//		500: ErrorMessage
func getParking(c echo.Context) error {
	parkingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error in parking id",
		})
	}

	parking, capacity, err := manager.GetParking(c.Request().Context(), parkingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"parking": toParkingRes(parking, capacity, uuid.UUID{})})
}

// swagger:route GET /v1/parkings System_Admin getParkings
//
// # This route is used to get all parkings
//
// responses:
//
//	200: ParkingsGetRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func getParkings(c echo.Context) error {
	parkings, err := manager.GetParkings(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"parkings": toParkingResSlice(parkings)})
}

// swagger:route PUT /v1/parking/{id} System_Admin updateParking
//
// # This route is used to update a parking
//
// responses:
//
//	201: ParkingUpdateRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func updateParking(c echo.Context) error {
	p := new(Parking)
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

	err = manager.UpdateParking(c.Request().Context(), p)
	if err != nil {
		if errors.Is(err, manager.ErrDuplicateEntity) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		if errors.Is(err, manager.ErrNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{"parking": toParkingRes(p, 0, uuid.UUID{})})
}

// swagger:route DELETE /v1/parking/{id} System_Admin deleteParking
//
// # This route is used to delete a parking by ID
//
// responses:
//
//	200: ErrorMessage
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func deleteParking(c echo.Context) error {
	parkingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "",
		})
	}
	err = manager.DeleteParking(c.Request().Context(), parkingID)
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
		"message": "Parking deleted successfully",
	})
}

// swagger:route GET /v1/user/parkings User getUserParkings
//
// # This route is used to get all user parkings which user can request access to them
//
// responses:
//
//	200: UserParkingsGetRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func getUserParkings(c echo.Context) error {
	user := c.Get("user").(*entity.CustomClaims)
	parkings, err := manager.GetUserParkings(c.Request().Context(), user.Phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"parkings": toParkingResSlice(parkings)})
}
