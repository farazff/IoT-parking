package rest

import (
	"github.com/farazff/IoT-parking/middleware"
	"github.com/okian/servo/v2/rest"
)

func init() {
	rest.EchoGet("/health_check", healthCheck)

	//rest.EchoPost("/v1/systemAdmin", createSystemAdmin, middleware.ApiKey)
	//rest.EchoGet("/v1/systemAdmin/:id", getSystemAdmin, middleware.ApiKey)
	//rest.EchoGet("/v1/systemAdmins", getSystemAdmins, middleware.ApiKey)
	//rest.EchoPut("/v1/systemAdmin/:id", updateSystemAdmin, middleware.ApiKey)
	//rest.EchoDelete("/v1/systemAdmin/:id", deleteSystemAdmin, middleware.ApiKey)

	rest.EchoPost("/v1/parking", createParking, middleware.AdminApiKey)
	rest.EchoGet("/v1/parking/:id", getParking, middleware.AdminApiKey)
	rest.EchoGet("/v1/parkings", getParkings, middleware.AdminApiKey)
	rest.EchoPut("/v1/parking/:id", updateParking, middleware.AdminApiKey)
	rest.EchoDelete("/v1/parking/:id", deleteParking, middleware.AdminApiKey)

	rest.EchoPost("/v1/parkingAdmin", createParkingAdmin, middleware.AdminApiKey)
	rest.EchoGet("/v1/parkingAdmin/:id", getParkingAdmin, middleware.AdminApiKey)
	rest.EchoGet("/v1/parkingAdmins", getParkingAdmins, middleware.AdminApiKey)
	rest.EchoPut("/v1/parkingAdmin/:id", updateParkingAdmin, middleware.AdminApiKey)
	rest.EchoDelete("/v1/parkingAdmin/:id", deleteParkingAdmin, middleware.AdminApiKey)

	rest.EchoPost("/v1/ParkingAdmin/signIn", parkingAdminSignIn)

	rest.EchoPost("/v1/zone", createZone, middleware.ApiKey)
	rest.EchoGet("/v1/zones", getZones, middleware.ApiKey)
	rest.EchoGet("/v1/zone/:id", getZone, middleware.ApiKey)
	rest.EchoPut("/v1/zone/:id", updateZone, middleware.ApiKey)
	rest.EchoDelete("/v1/zone/:id", deleteZone, middleware.ApiKey)
	rest.EchoPut("/v1/zoneEnter/:id/:uuid", EnterZone, middleware.ApiKey)
	rest.EchoPut("/v1/zoneExit/:id/:uuid", ExitZone, middleware.ApiKey)

	rest.EchoPost("/v1/whitelist", createWhitelist, middleware.ApiKey)
	rest.EchoGet("/v1/whitelists", getWhitelists, middleware.ApiKey)
	rest.EchoDelete("/v1/whitelist/:id", deleteWhitelist, middleware.ApiKey)

	rest.EchoPost("v1/carEnter/:uuid", carEnter, middleware.ApiKey)
	rest.EchoPut("v1/carExit/:uuid", carExit, middleware.ApiKey)
}
