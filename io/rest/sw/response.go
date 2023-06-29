//		Parking IoT Swagger
//
//		Document of Parking-IoT api
//
//		Host: localhost:7676
//	    Schemes: http
//	    BasePath: /
//	    Version: 1.0.0
//	    Contact: Faraz Farangizadeh<f.farangizadeh@gmail.com>
//
//	    Consumes:
//	    - application/json
//
//	    Produces:
//	    - application/json
//
// swagger:meta
package sw

// swagger:response NoContent
type NoContent struct {
	// in: body
	Body struct{}
	//in: cookie
	SessionToken string
}

// swagger:response ErrorUnauthorizedMessage
type ErrorUnauthorizedMessage struct {
	// in: body
	Body struct {
		// Error message
		Message string `json:"message"`
	}
}

// swagger:response ErrorMessage
type ErrorMessage struct {
	// in: body
	Body struct {
		// Error message
		Message string `json:"message"`
	}
	//in: cookie
	SessionToken string
}

// swagger:model
type ParkingCreate struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Enabled bool   `json:"enabled"`
	Uuid    string `json:"uuid"`
}

// swagger:response ParkingCreateRes
type ParkingCreateRes struct {
	//in: body
	Body struct {
		// parking res
		ParkingCreate ParkingCreate `json:"parking"`
	}
	//in: cookie
	SessionToken string
}

// swagger:model
type ParkingGet struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Enabled  bool   `json:"enabled"`
	Capacity int    `json:"capacity"`
}

// swagger:response ParkingGetRes
type ParkingGetRes struct {
	//in: body
	Body struct {
		// parking res
		ParkingGet ParkingGet `json:"parking"`
	}
	//in: cookie
	SessionToken string
}

// swagger:model
type ParkingsGet struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Enabled bool   `json:"enabled"`
}

// swagger:response ParkingsGetRes
type ParkingsGetRes struct {
	//in: body
	Body struct {
		// parking res
		ParkingsGet []ParkingsGet `json:"parkings"`
	}
	//in: cookie
	SessionToken string
}

// swagger:response ParkingUpdateRes
type ParkingsUpdateRes struct {
	//in: body
	Body struct {
		// parking res
		ParkingUpdate []ParkingsGet `json:"parking"`
	}
	//in: cookie
	SessionToken string
}
