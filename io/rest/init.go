package rest

import (
	"github.com/farazff/IoT-parking/middleware"
	"github.com/okian/servo/v2/rest"
)

func init() {
	docs()
	rest.EchoGet("/health_check", healthCheck)

	//rest.EchoPost("/v1/systemAdmin", createSystemAdmin, middleware.ApiKey)
	//rest.EchoGet("/v1/systemAdmin/:id", getSystemAdmin, middleware.ApiKey)
	//rest.EchoGet("/v1/systemAdmins", getSystemAdmins, middleware.ApiKey)
	//rest.EchoPut("/v1/systemAdmin/:id", updateSystemAdmin, middleware.ApiKey)
	//rest.EchoDelete("/v1/systemAdmin/:id", deleteSystemAdmin, middleware.ApiKey)

	//System admin requests
	rest.EchoPost("/v1/systemAdmin/signIn", systemAdminSignIn, middleware.SystemAdminApiKey)

	rest.EchoPost("/v1/parking", createParking, middleware.SystemAdminApiKey)
	rest.EchoGet("/v1/parking/:id", getParking, middleware.SystemAdminApiKey)
	rest.EchoGet("/v1/parkings", getParkings, middleware.SystemAdminApiKey)
	rest.EchoPut("/v1/parking/:id", updateParking, middleware.SystemAdminApiKey)
	rest.EchoDelete("/v1/parking/:id", deleteParking, middleware.SystemAdminApiKey)

	rest.EchoPost("/v1/parkingAdmin", createParkingAdmin, middleware.SystemAdminApiKey)
	rest.EchoGet("/v1/parkingAdmin/:id", getParkingAdmin, middleware.SystemAdminApiKey)
	rest.EchoGet("/v1/parkingAdmins", getParkingAdmins, middleware.SystemAdminApiKey)
	rest.EchoPut("/v1/parkingAdmin/:id", updateParkingAdmin, middleware.SystemAdminApiKey)
	rest.EchoDelete("/v1/parkingAdmin/:id", deleteParkingAdmin, middleware.SystemAdminApiKey)

	//Parking admin requests
	rest.EchoPost("/v1/parkingAdmin/signIn", parkingAdminSignIn, middleware.ParkingAdminApiKey)

	rest.EchoPost("/v1/zone", createZone, middleware.ParkingAdminApiKey)
	rest.EchoGet("/v1/zones", getZones, middleware.ParkingAdminApiKey)
	rest.EchoGet("/v1/zone/:id", getZone, middleware.ParkingAdminApiKey)
	rest.EchoPut("/v1/zone/:id", updateZone, middleware.ParkingAdminApiKey)
	rest.EchoDelete("/v1/zone/:id", deleteZone, middleware.ParkingAdminApiKey)

	rest.EchoPut("/v1/whitelist/approve/:id", approveWhitelist, middleware.ParkingAdminApiKey)
	rest.EchoGet("/v1/whitelists/approved", getWhitelistsApproved, middleware.ParkingAdminApiKey)
	rest.EchoGet("/v1/whitelists/toApprove", getWhitelistsToApprove, middleware.ParkingAdminApiKey)
	rest.EchoDelete("/v1/whitelist/:id", deleteWhitelist, middleware.ParkingAdminApiKey)

	//Raspberry PI requests
	rest.EchoPost("v1/carEnter/:uuid/:tag", carEnter, middleware.HardwareApiKey)
	rest.EchoPut("v1/carExit/:uuid/:tag", carExit, middleware.HardwareApiKey)

	rest.EchoPut("/v1/zoneEnter/:id/:uuid", enterZone, middleware.HardwareApiKey)
	rest.EchoPut("/v1/zoneExit/:id/:uuid", exitZone, middleware.HardwareApiKey)

	//User requests
	rest.EchoPost("/v1/user/signIn", userSignIn, middleware.UserApiKey)

	rest.EchoPost("/v1/user/signUp", userSignUp, middleware.UserApiKey)

	rest.EchoPost("/v1/user/whitelist/request", requestWhitelist, middleware.UserApiKey)
	rest.EchoGet("/v1/user/whitelists/approved", getUserWhitelists, middleware.UserApiKey)

	rest.EchoGet("/v1/user/logs/:page", getUserLogs, middleware.UserApiKey)

	rest.EchoGet("/v1/user/parkings", getUserParkings, middleware.UserApiKey)
}
