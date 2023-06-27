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

	//System admin requests
	rest.EchoPost("/v1/systemAdmin/signIn", systemAdminSignIn)

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

	//Parking admin requests
	rest.EchoPost("/v1/parkingAdmin/signIn", parkingAdminSignIn)

	rest.EchoPost("/v1/zone", createZone, middleware.ApiKey)
	rest.EchoGet("/v1/zones", getZones, middleware.ApiKey)
	rest.EchoGet("/v1/zone/:id", getZone, middleware.ApiKey)
	rest.EchoPut("/v1/zone/:id", updateZone, middleware.ApiKey)
	rest.EchoDelete("/v1/zone/:id", deleteZone, middleware.ApiKey)

	rest.EchoPut("/v1/whitelist/approve/:id", approveWhitelist, middleware.ApiKey)
	rest.EchoGet("/v1/whitelists/approved", getWhitelistsApproved, middleware.ApiKey)
	rest.EchoGet("/v1/whitelists/toApprove", getWhitelistsToApprove, middleware.ApiKey)
	rest.EchoDelete("/v1/whitelist/:id", deleteWhitelist, middleware.ApiKey)

	//Raspberry PI requests
	rest.EchoPost("v1/carEnter/:uuid", carEnter, middleware.ApiKey)
	rest.EchoPut("v1/carExit/:uuid", carExit, middleware.ApiKey)

	rest.EchoPut("/v1/zoneEnter/:id/:uuid", enterZone, middleware.ApiKey)
	rest.EchoPut("/v1/zoneExit/:id/:uuid", exitZone, middleware.ApiKey)

	//User requests
	rest.EchoPost("/v1/user/signIn", userSignIn)

	rest.EchoPost("/v1/white/request", requestWhitelist)

}
