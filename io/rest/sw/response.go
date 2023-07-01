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

import "time"

// swagger:response NoContent
type NoContent struct {
	// in: body
	Body struct{}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:response ErrorUnauthorizedMessage
type ErrorUnauthorizedMessage struct {
	// in: body
	Body struct {
		// Error message
		Message string `json:"message"`
	}
}

// swagger:response ErrorMessage JustMessage
type ErrorMessage struct {
	// in: body
	Body struct {
		// Error message
		Message string `json:"message"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
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
	SessionToken string `json:"session_token"`
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
	SessionToken string `json:"session_token"`
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
	SessionToken string `json:"session_token"`
}

// swagger:response ParkingUpdateRes
type ParkingUpdateRes struct {
	//in: body
	Body struct {
		// parking res
		ParkingUpdate ParkingGet `json:"parking"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:model
type ParkingAdminCreate struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Enabled   bool   `json:"enabled"`
	ParkingID int    `json:"parking_id"`
}

// swagger:response ParkingAdminCreateRes
type ParkingAdminCreateRes struct {
	//in: body
	Body struct {
		// parking admin res
		ParkingAdminCreate ParkingAdminCreate `json:"parking_admin"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:model
type ParkingAdminGet struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Enabled   bool   `json:"enabled"`
	ParkingID int    `json:"parking_id"`
}

// swagger:response ParkingAdminGetRes
type ParkingAdminGetRes struct {
	//in: body
	Body struct {
		// parking res
		ParkingAdminGet ParkingAdminGet `json:"parking_admin"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:response ParkingAdminsGetRes
type ParkingAdminsGetRes struct {
	//in: body
	Body struct {
		// parking res
		ParkingAdminGet []ParkingAdminGet `json:"parking_admins"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:response ParkingAdminUpdateRes
type ParkingAdminUpdateRes struct {
	//in: body
	Body struct {
		// parking res
		ParkingAdminGet ParkingAdminGet `json:"parking_admin"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:model
type ZoneCreate struct {
	ID               int  `json:"id"`
	Capacity         int  `json:"capacity"`
	Enabled          bool `json:"enabled"`
	RemainedCapacity int  `json:"remained_capacity"`
	ParkingID        int  `json:"parking_id"`
}

// swagger:response ZoneCreateRes
type ZoneCreateRes struct {
	//in: body
	Body struct {
		// zone res
		ZoneCreate ZoneCreate `json:"zone"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:model
type ZoneGet struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Enabled   bool   `json:"enabled"`
	ParkingID int    `json:"parking_id"`
}

// swagger:response ParkingAdminGetRes
type ZoneGetRes struct {
	//in: body
	Body struct {
		// parking res
		ZoneGet ZoneGet `json:"zone"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:response ZonesGetRes
type ZonesGetRes struct {
	//in: body
	Body struct {
		// parking res
		ZonesGet []ZoneGet `json:"zones"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:model
type UserParkingsGet struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Enabled    bool   `json:"enabled"`
	HaveAccess bool   `json:"have_access"`
}

// swagger:response UserParkingsGetRes
type UserParkingsGetRes struct {
	//in: body
	Body struct {
		// parking res
		UserParkingsGet []UserParkingsGet `json:"parkings"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:model
type UserRequestsGet struct {
	ID             int    `json:"id"`
	ParkingName    string `json:"parking_name"`
	ParkingAddress string `json:"parking_address"`
	Approved       string `json:"approved"`
}

// swagger:response UserRequestsRes
type UserRequestsRes struct {
	//in: body
	Body struct {
		// parking res
		UserRequestsGet []UserRequestsGet `json:"whitelists"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}

// swagger:model
type UserLogsGet struct {
	ID             int       `json:"id"`
	EnterTime      time.Time `json:"enter_time"`
	ExitTime       time.Time `json:"exit_time"`
	ParkingName    string    `json:"parking_name"`
	ParkingAddress string    `json:"parking_address"`
}

// swagger:response UserLogsRes
type UserLogsRes struct {
	//in: body
	Body struct {
		// parking res
		UserLogsGet []UserLogsGet `json:"logs"`
	}
	//in: cookie
	SessionToken string `json:"session_token"`
}
