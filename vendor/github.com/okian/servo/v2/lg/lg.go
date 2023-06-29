package lg

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/spf13/viper"
)

func (s *service) Info(args ...interface{}) {
	s.logger.Info(args...)
}

func (s *service) Debug(args ...interface{}) {
	if viper.GetBool("lg_dump") {
		spew.Dump(args...)
		return
	}
	s.logger.Debug(args...)
}

func (s *service) Warn(args ...interface{}) {
	s.logger.Warn(args...)
}

func (s *service) Error(args ...interface{}) {
	s.logger.Error(args...)
}

func (s *service) Panic(args ...interface{}) {
	s.logger.Panic(args...)
}

func (s *service) Fatal(args ...interface{}) {
	s.logger.Fatal(args...)
}

func (s *service) Infof(template string, args ...interface{}) {
	s.logger.Infof(template, args...)
}

func (s *service) Debugf(template string, args ...interface{}) {
	s.logger.Debugf(template, args...)
}

func (s *service) Warnf(template string, args ...interface{}) {
	s.logger.Warnf(template, args...)
}

func (s *service) Errorf(template string, args ...interface{}) {
	s.logger.Errorf(template, args...)
}

func (s *service) Panicf(template string, args ...interface{}) {
	s.logger.Panicf(template, args...)
}

func (s *service) Fatalf(template string, args ...interface{}) {
	s.logger.Fatalf(template, args...)
}

func (s *service) Infow(template string, keysAndValues ...interface{}) {
	s.logger.Infow(template, keysAndValues...)
}

func (s *service) Debugw(template string, keysAndValues ...interface{}) {
	s.logger.Debugw(template, keysAndValues...)
}

func (s *service) Warnw(template string, keysAndValues ...interface{}) {
	s.logger.Warnw(template, keysAndValues...)
}

func (s *service) Errorw(template string, keysAndValues ...interface{}) {
	s.logger.Errorw(template, keysAndValues...)
}

func (s *service) Errorwt(ctx context.Context, msg string, keysAndValues ...interface{}) {
	trace(ctx, msg, keysAndValues...)
	s.logger.Errorw(msg, keysAndValues...)
}

func (s *service) Panicw(template string, keysAndValues ...interface{}) {
	s.logger.Panicw(template, keysAndValues...)
}

func (s *service) Fatalw(template string, keysAndValues ...interface{}) {
	s.logger.Fatalw(template, keysAndValues...)
}

func trace(ctx context.Context, msg string, keysAndValues ...interface{}) {
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		return
	}
	res := make([]log.Field, 0)
	if msg != "" {
		res = append(res, log.String("msg", msg))
	}
	for i := 0; i < len(keysAndValues); {
		switch keysAndValues[i+1].(type) {
		case string:
			res = append(res, log.String(keysAndValues[i].(string), keysAndValues[i+1].(string)))
		case int:
			res = append(res, log.Int(keysAndValues[i].(string), keysAndValues[i+1].(int)))
		case int8:
			res = append(res, log.Int(keysAndValues[i].(string), int(keysAndValues[i+1].(int8))))
		case int16:
			res = append(res, log.Int(keysAndValues[i].(string), int(keysAndValues[i+1].(int16))))
		case int32:
			res = append(res, log.Int32(keysAndValues[i].(string), keysAndValues[i+1].(int32)))
		case int64:
			res = append(res, log.Int64(keysAndValues[i].(string), keysAndValues[i+1].(int64)))
		case float32:
			res = append(res, log.Float32(keysAndValues[i].(string), keysAndValues[i+1].(float32)))
		case float64:
			res = append(res, log.Float64(keysAndValues[i].(string), keysAndValues[i+1].(float64)))
		case bool:
			res = append(res, log.Bool(keysAndValues[i].(string), keysAndValues[i+1].(bool)))
		default:
			res = append(res, log.Object(keysAndValues[i].(string), keysAndValues[i+1]))
		}
		i += 2
	}
	sp.LogFields(res...)
}
