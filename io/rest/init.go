package rest

import (
	"github.com/okian/servo/v2/rest"
)

func init() {
	rest.EchoPost("/v1/parking", createParking, apiKey)
	rest.EchoGet("/v1/parking/:id", getParking, apiKey)
	rest.EchoGet("/v1/parkings", getParkings, apiKey)
	rest.EchoPut("/v1/parking/:id", updateParking, apiKey)
	rest.EchoDelete("/v1/parking/:id", deleteParking, apiKey)

	rest.EchoPost("/v1/zone", createZone, apiKey)
	rest.EchoGet("/v1/zone/:id", getZone, apiKey)
	rest.EchoGet("/v1/zones", getZones, apiKey)
	rest.EchoPut("/v1/zone/:id", updateZone, apiKey)
	rest.EchoDelete("/v1/zone/:id", deleteZone, apiKey)

	rest.EchoPost("/v1/parkingAdmin", createParkingAdmin, apiKey)
	rest.EchoGet("/v1/parkingAdmin/:id", getParkingAdmin, apiKey)
	rest.EchoGet("/v1/parkingAdmins", getParkingAdmin, apiKey)
	rest.EchoPut("/v1/parkingAdmin/:id", updateParkingAdmin, apiKey)
	rest.EchoDelete("/v1/parkingAdmin/:id", deleteParkingAdmin, apiKey)

	rest.EchoPost("/v1/systemAdmin", createSystemAdmin, apiKey)
	rest.EchoGet("/v1/systemAdmin/:id", getSystemAdmin, apiKey)
	rest.EchoGet("/v1/systemAdmins", getsyStemAdmin, apiKey)
	rest.EchoPut("/v1/systemAdmin/:id", updateSystemAdmin, apiKey)
	rest.EchoDelete("/v1/systemAdmin/:id", deleteSystemAdmin, apiKey)
}
