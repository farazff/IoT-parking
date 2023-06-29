package config

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/okian/servo/v2"
	"github.com/spf13/viper"
)

const configFile = "config"

type cfg struct{}

func (c *cfg) Name() string {
	return "config"
}

func (c *cfg) Initialize(_ context.Context) error {
	viper.SetEnvPrefix(AppName())
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", AppName()))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s/", AppName()))
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	if viper.GetString("tz") != "" {
		z, err := time.LoadLocation(viper.GetString("tz"))
		if err != nil {
			return err
		}
		time.Local = z

	}
	return nil
}

func (c *cfg) Finalize() error {
	return nil
}

func init() {
	c := &cfg{}
	servo.Register(c, 10)
}
