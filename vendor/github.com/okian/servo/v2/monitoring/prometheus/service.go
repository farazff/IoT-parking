package prometheus

import (
	"net/http"

	"github.com/spf13/viper"
)

type service struct {
	server *http.Server
}

func Namespace() string {
	return viper.GetString("monitoring_namespace")
}
