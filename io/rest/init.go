package rest

import (
	"github.com/farazff/IoT-parking/middleware"
	"github.com/okian/servo/v2/rest"
)

func init() {
	rest.EchoGet("/health_check", healthCheck)

	rest.EchoPost("/v1/parking", createParking, middleware.ApiKey)
	rest.EchoGet("/v1/parking/:id", getParking, middleware.ApiKey)
	rest.EchoGet("/v1/parkings", getParkings, middleware.ApiKey)
	rest.EchoPut("/v1/parking/:id", updateParking, middleware.ApiKey)
	rest.EchoDelete("/v1/parking/:id", deleteParking, middleware.ApiKey)

	rest.EchoPost("/v1/zone", createZone, middleware.ApiKey)
	rest.EchoGet("/v1/zone/:id", getZone, middleware.ApiKey)
	rest.EchoGet("/v1/zones", getZones, middleware.ApiKey)
	rest.EchoPut("/v1/zone/:id", updateZone, middleware.ApiKey)
	rest.EchoDelete("/v1/zone/:id", deleteZone, middleware.ApiKey)

	rest.EchoPost("/v1/parkingAdmin", createParkingAdmin, middleware.ApiKey)
	rest.EchoGet("/v1/parkingAdmin/:id", getParkingAdmin, middleware.ApiKey)
	rest.EchoGet("/v1/parkingAdmins", getParkingAdmin, middleware.ApiKey)
	rest.EchoPut("/v1/parkingAdmin/:id", updateParkingAdmin, middleware.ApiKey)
	rest.EchoDelete("/v1/parkingAdmin/:id", deleteParkingAdmin, middleware.ApiKey)

	rest.EchoPost("/v1/systemAdmin", createSystemAdmin, middleware.ApiKey)
	rest.EchoGet("/v1/systemAdmin/:id", getSystemAdmin, middleware.ApiKey)
	rest.EchoGet("/v1/systemAdmins", getSystemAdmins, middleware.ApiKey)
	rest.EchoPut("/v1/systemAdmin/:id", updateSystemAdmin, middleware.ApiKey)
	rest.EchoDelete("/v1/systemAdmin/:id", deleteSystemAdmin, middleware.ApiKey)

	rest.EchoPost("/v1/whitelist", createWhitelist, middleware.ApiKey)
	rest.EchoGet("/v1/whitelists", getWhitelists, middleware.ApiKey)
	rest.EchoDelete("/v1/WhiteList", deleteWhitelist, middleware.ApiKey)

	rest.EchoPost("v1/carEnter", carEnter, middleware.ApiKey)
	rest.EchoPut("v1/carExit", carExit, middleware.ApiKey)
}
