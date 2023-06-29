package lg

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func fileWriter() (zapcore.WriteSyncer, error) {
	path := viper.GetString("log_filepath")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(f), nil
}
