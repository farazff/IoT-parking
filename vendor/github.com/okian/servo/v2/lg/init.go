package lg

import (
	"github.com/okian/servo/v2"
)

func init() {
	s := &service{}
	servo.Register(s, 20)
}
