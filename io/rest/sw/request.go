package sw

// swagger:parameters
type apiKey struct {
	// api key for APIs
	// in: header
	// required: true
	ApiKey string `json:"api-key"`
}

// swagger:parameters parkingAdminSingIn systemAdminSignIn userSingIn
type signIN struct {
	apiKey
	// in: body
	// required: true
	Phone string `json:"phone"`
	// in: body
	// required: true
	Password string `json:"password"`
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
	// in: body
	// required: true
	User User
}
