package rest

import (
	"github.com/farazff/IoT-parking/middleware"
	cors "github.com/labstack/echo/v4/middleware"
	"github.com/okian/servo/v2/rest"
)

func init() {
	customCORS := cors.CORSWithConfig(cors.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Set-Cookie", "session_token"},
	})

	rest.Use(customCORS)

	docs()

	rest.EchoGet("/health_check", healthCheck, customCORS)

	//rest.EchoPost("/v1/systemAdmin", createSystemAdmin, middleware.ApiKey)
	//rest.EchoGet("/v1/systemAdmin/:id", getSystemAdmin, middleware.ApiKey)
	//rest.EchoGet("/v1/systemAdmins", getSystemAdmins, middleware.ApiKey)
	//rest.EchoPut("/v1/systemAdmin/:id", updateSystemAdmin, middleware.ApiKey)
	//rest.EchoDelete("/v1/systemAdmin/:id", deleteSystemAdmin, middleware.ApiKey)

	//System admin requests
	rest.EchoPost("/systemAdmin/signIn", systemAdminSignIn, middleware.SystemAdminApiKey, customCORS)
	rest.EchoPost("/systemAdmin/refresh-token", refreshToken, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemTokenExpired)
	rest.EchoPost("/systemAdmin/signOut", systemAdminSignOut, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)

	rest.EchoPost("/v1/parking", createParking, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)
	rest.EchoGet("/v1/parking/:id", getParking, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)
	rest.EchoGet("/v1/parkings", getParkings, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)
	rest.EchoPut("/v1/parking/:id", updateParking, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)
	rest.EchoDelete("/v1/parking/:id", deleteParking, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)

	rest.EchoPost("/v1/parkingAdmin", createParkingAdmin, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)
	rest.EchoGet("/v1/parkingAdmin/:id", getParkingAdmin, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)
	rest.EchoGet("/v1/parkingAdmins", getParkingAdmins, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)
	rest.EchoPut("/v1/parkingAdmin/:id", updateParkingAdmin, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)
	rest.EchoDelete("/v1/parkingAdmin/:id", deleteParkingAdmin, middleware.SystemAdminApiKey, customCORS, middleware.ValidateSystemToken)

	//Parking admin requests
	rest.EchoPost("/parkingAdmin/signIn", parkingAdminSignIn, middleware.ParkingAdminApiKey, customCORS)
	rest.EchoPost("/parkingAdmin/refresh-token", refreshToken, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingTokenExpired)
	rest.EchoPost("/parkingAdmin/signOut", parkingAdminSignOut, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)

	rest.EchoPost("/v1/zone", createZone, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)
	rest.EchoGet("/v1/zones", getZones, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)
	rest.EchoGet("/v1/zone/:id", getZone, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)
	rest.EchoPut("/v1/zone/:id", updateZone, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)
	rest.EchoDelete("/v1/zone/:id", deleteZone, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)

	rest.EchoGet("/v1/whitelists/toApprove", getWhitelistsToApprove, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)
	rest.EchoPut("/v1/whitelist/approve/:id", approveWhitelist, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)
	rest.EchoGet("/v1/whitelists/approved", getWhitelistsApproved, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)
	rest.EchoDelete("/v1/whitelist/:id", deleteWhitelist, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)

	rest.EchoGet("/v1/parking", getAdminParking, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)

	rest.EchoGet("/v1/logs/:page", getLogs, middleware.ParkingAdminApiKey, customCORS, middleware.ValidateParkingToken)

	//User requests
	rest.EchoPost("/user/signUp", userSignUp, middleware.UserApiKey, customCORS)
	rest.EchoPost("/user/signIn", userSignIn, middleware.UserApiKey, customCORS)
	rest.EchoPost("/user/refresh-token", refreshToken, middleware.UserApiKey, customCORS, middleware.ValidateUserTokenExpired)
	rest.EchoPost("/user/signOut", userSignOut, middleware.UserApiKey, customCORS, middleware.ValidateUserToken)

	rest.EchoGet("/v1/user/parkings", getUserParkings, middleware.UserApiKey, customCORS, middleware.ValidateUserToken)
	rest.EchoGet("/v1/user/whitelists/requested", getUserWhitelists, middleware.UserApiKey, customCORS, middleware.ValidateUserToken)
	rest.EchoPost("/v1/user/whitelist/request", requestWhitelist, middleware.UserApiKey, customCORS, middleware.ValidateUserToken)

	rest.EchoGet("/v1/user/logs/:page", getUserLogs, middleware.UserApiKey, customCORS, middleware.ValidateUserToken)

	rest.EchoGet("/v1/user", getUser, middleware.UserApiKey, customCORS, middleware.ValidateUserToken)
	rest.EchoPut("/v1/user", updateUser, middleware.UserApiKey, customCORS, middleware.ValidateUserToken)

	//Raspberry PI requests
	rest.EchoPost("v1/carEnter/:uuid/:tag", carEnter, middleware.HardwareApiKey)
	rest.EchoPut("v1/carExit/:uuid/:tag", carExit, middleware.HardwareApiKey)

	rest.EchoPut("/v1/zoneEnter/:id/:uuid", enterZone, middleware.HardwareApiKey)
	rest.EchoPut("/v1/zoneExit/:id/:uuid", exitZone, middleware.HardwareApiKey)
}
