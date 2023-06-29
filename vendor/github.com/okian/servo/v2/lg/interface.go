package lg

import "context"

type Interface interface {
	Name() string
	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Infof(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Errorwt(ctx context.Context, msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

var (
	logger Interface
)

func Register(i Interface) {
	if logger != nil {
		panic("multiple call")
	}
	logger = i
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func Infow(template string, keysAndValues ...interface{}) {
	logger.Infow(template, keysAndValues...)
}

func Debugw(template string, keysAndValues ...interface{}) {
	logger.Debugw(template, keysAndValues...)
}

func Warnw(template string, keysAndValues ...interface{}) {
	logger.Warnw(template, keysAndValues...)
}

func Errorw(template string, keysAndValues ...interface{}) {
	logger.Errorw(template, keysAndValues...)
}

func Panicw(template string, keysAndValues ...interface{}) {
	logger.Panicw(template, keysAndValues...)
}

func Fatalw(template string, keysAndValues ...interface{}) {
	logger.Fatalw(template, keysAndValues...)
}
