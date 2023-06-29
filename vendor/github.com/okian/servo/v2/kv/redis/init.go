package redis

import (
	"github.com/okian/servo/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault(host, "127.0.0.1")
	viper.SetDefault(port, 6379)
	servo.Register(&service{}, 20)
}
