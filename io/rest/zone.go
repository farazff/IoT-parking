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

// swagger:route POST /v1/zone Parking_Admin createZone
//
// # This route is used to create zone
//
// responses:
//
//	201: ZoneCreateRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func createZone(c echo.Context) error {
	phone := c.Get("user").(*entity.CustomClaims).Phone

	z := new(Zone)
	if err := c.Bind(z); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(z); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	id, parkingID, err := manager.CreateZone(c.Request().Context(), z, phone)
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
	z.FParkingID = parkingID
	return c.JSON(http.StatusCreated, echo.Map{"zone": toZoneRes(z, id)})
}

// swagger:route GET /v1/zones Parking_Admin getZones
//
// # This route is used to get all zones
//
// responses:
//
//	200: ZonesGetRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	500: ErrorMessage
func getZones(c echo.Context) error {
	phone := c.Get("user").(*entity.CustomClaims).Phone

	zones, err := manager.GetZones(c.Request().Context(), phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{"zones": toZoneResSlice(zones)})
}

// swagger:route GET /v1/zone/{id} Parking_Admin getZone
//
// # This route is used to get a single zone by ID
//
// responses:
//
//	200: ZoneGetRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func getZone(c echo.Context) error {
	phone := c.Get("user").(*entity.CustomClaims).Phone

	zoneID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	zone, err := manager.GetZone(c.Request().Context(), int(zoneID), phone)
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
	return c.JSON(http.StatusOK, echo.Map{"zone": toZoneRes(zone, -1)})
}

// swagger:route PUT /v1/zone/{id} Parking_Admin updateZone
//
// # This route is used to update a zone
//
// responses:
//
//	201: ZoneUpdateRes
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func updateZone(c echo.Context) error {
	phone := c.Get("user").(*entity.CustomClaims).Phone

	z := new(Zone)
	if err := c.Bind(z); err != nil {
		lg.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := c.Validate(z); err != nil {
		lg.Error("body validation failed")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "body validation failed",
		})
	}

	zoneID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	z.FID = int(zoneID)

	err = manager.UpdateZone(c.Request().Context(), z, phone)
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

	return c.JSON(http.StatusCreated, echo.Map{"zone": toZoneRes(z, -1)})
}

// swagger:route DELETE /v1/zone/{id} Parking_Admin deleteZone
//
// # This route is used to delete a zone by ID
//
// responses:
//
//	200: ErrorMessage
//	400: ErrorMessage
//	401: ErrorUnauthorizedMessage
//	404: ErrorMessage
//	500: ErrorMessage
func deleteZone(c echo.Context) error {
	phone := c.Get("user").(*entity.CustomClaims).Phone

	ZoneID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "",
		})
	}

	err = manager.DeleteZone(c.Request().Context(), ZoneID, phone)
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
		"message": "Zone deleted successfully",
	})
}

func enterZone(c echo.Context) error {
	zid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	parkingUUID := c.Param("uuid")
	_, err = uuid.Parse(parkingUUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.EnterZone(c.Request().Context(), int(zid), parkingUUID)
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

	return c.JSON(http.StatusCreated, echo.Map{"message": "Updated successfully"})
}

func exitZone(c echo.Context) error {
	zid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	parkingUUID := c.Param("uuid")
	_, err = uuid.Parse(parkingUUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = manager.ExitZone(c.Request().Context(), int(zid), parkingUUID)
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

	return c.JSON(http.StatusCreated, echo.Map{"message": "Updated successfully"})
}
