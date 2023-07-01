package sw

// swagger:model
type apiKey struct {
	// api key for APIs
	// in: header
	// required: true
	ApiKey string `json:"api-key"`
}

// swagger:model
type sessionToken struct {
	// Session token
	// in: header
	// required: true
	SessionToken string `json:"session_token"`
}

// swagger:model
type Credentials struct {
	// in: body
	// required: true
	Phone string `json:"phone"`
	// in: body
	// required: true
	Password string `json:"password"`
}

// swagger:parameters parkingAdminSingIn systemAdminSignIn userSingIn
type signIN struct {
	apiKey
	// required: true
	// in: body
	Credentials Credentials
}

// swagger:model
type User struct {
	// required: true
	FirstName string `json:"first_name"`
	// required: true
	LastName string `json:"last_name"`
	// required: true
	CarTag string `json:"car_tag"`
	// required: true
	Phone string `json:"phone"`
	// required: true
	Password string `json:"password"`
}

// swagger:parameters userSingUp
type signUp struct {
	apiKey
	// required: true
	// in: body
	User User
}

// swagger:model
type ParkingCreator struct {
	// required: true
	Name string `json:"name"`
	// required: true
	Address string `json:"address"`
	// required: true
	Phone string `json:"phone"`
	// required: true
	Enabled bool `json:"enabled"`
}

// swagger:parameters createParking
type createParking struct {
	apiKey
	sessionToken
	// required: true
	// in: body
	ParkingCreator ParkingCreator
}

// swagger:parameters getParking deleteParking
type getParking struct {
	apiKey
	sessionToken
	// required: true
	// in: path
	ID int
}

// swagger:parameters getParkings
type getParkings struct {
	apiKey
	sessionToken
}

// swagger:parameters updateParking
type updateParking struct {
	apiKey
	sessionToken
	// required: true
	// in: path
	ID int
	// required: true
	// in: body
	ParkingCreator ParkingCreator `json:"parkingUpdater"`
}

// swagger:model
type ParkingAdminCreator struct {
	// required: true
	FirstName string `json:"first_name"`
	// required: true
	LastName string `json:"last_name"`
	// required: true
	Phone string `json:"phone"`
	// required: true
	Enabled bool `json:"enabled"`
	// required: true
	Password string `json:"password"`
	// required: true
	ParkingID int `json:"parking_id"`
}

// swagger:parameters createParkingAdmin
type createParkingAdmin struct {
	apiKey
	sessionToken
	// required: true
	// in: body
	ParkingAdminCreator ParkingAdminCreator
}

// swagger:parameters getParkingAdmin deleteParkingAdmin
type getParkingAdmin struct {
	apiKey
	sessionToken
	// required: true
	// in: path
	ID int
}

// swagger:parameters getParkingAdmins
type getParkingAdmins struct {
	apiKey
	sessionToken
}

// swagger:model
type ParkingAdminUpdater struct {
	// required: true
	FirstName string `json:"first_name"`
	// required: true
	LastName string `json:"last_name"`
	// required: true
	Phone string `json:"phone"`
	// required: true
	Enabled bool `json:"enabled"`
	// required: true
	ParkingID int `json:"parking_id"`
}

// swagger:parameters updateParkingAdmin
type updateParkingAdmin struct {
	apiKey
	sessionToken
	// required: true
	// in: path
	ID int
	// required: true
	// in: body
	ParkingAdminUpdater ParkingAdminUpdater `json:"parkingUpdater"`
}

// swagger:model
type ZoneCreator struct {
	// required: true
	Capacity string `json:"capacity"`
	// required: true
	Enabled string `json:"enabled"`
	// required: true
	RemainedCapacity string `json:"remained_capacity"`
}

// swagger:parameters createZone
type createZone struct {
	apiKey
	sessionToken
	// required: true
	// in: body
	ZoneCreator ZoneCreator
}

// swagger:parameters getZone deleteZone
type getZone struct {
	apiKey
	sessionToken
	// required: true
	// in: path
	ID int
}

// swagger:parameters getZones
type getZones struct {
	apiKey
	sessionToken
}

// swagger:parameters updateParkingAdmin
type updateZone struct {
	apiKey
	sessionToken
	// required: true
	// in: path
	ID int
	// required: true
	// in: body
	ZoneCreator ZoneCreator `json:"zoneUpdater"`
}

// swagger:parameters getUserParkings getUserRequests getUserLogs
type getUserParkings struct {
	apiKey
	sessionToken
}

// swagger:model
type accessRequestCreator struct {
	// required: true
	ParkingID string `json:"parking_id"`
}

// swagger:parameters accessRequest
type accessRequest struct {
	apiKey
	sessionToken
	// required: true
	// in: body
	accessRequestCreator accessRequestCreator
}
