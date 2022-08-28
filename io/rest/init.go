package rest

import (
	"github.com/okian/servo/v2/rest"
)

func init() {
	rest.EchoGet("/info", hello)
}
