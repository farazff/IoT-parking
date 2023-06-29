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
	SessionToken int
}

// swagger:response ErrorMessage
type ErrorMessage struct {
	// in: body
	Body struct {
		// Error message
		Message string `json:"message"`
	}
}
