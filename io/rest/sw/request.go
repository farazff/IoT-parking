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
