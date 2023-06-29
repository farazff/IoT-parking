package lg

import (
	"log/syslog"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func newSysLog() (zapcore.WriteSyncer, error) {
	w, err := syslog.Dial(viper.GetString("log_syslog_network"),
		viper.GetString("log_syslog_address"),
		syslog.LOG_INFO,
		"")
	return zapcore.AddSync(w), err
}
