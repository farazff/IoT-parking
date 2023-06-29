package kv

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type Interface interface {
	Name() string
	BitSet(ctx context.Context, key string, idx int, val bool, ttl time.Duration) error
	BitSets(ctx context.Context, key string, val bool, ttl time.Duration, idx ...int) error
	BitGet(ctx context.Context, key string, idx int) (bool, error)
	Set(ctx context.Context, key string, val string, ttl time.Duration) error
	Get(ctx context.Context, key string, rcv *string) error
	MSet(ctx context.Context, key string, val interface{}, ttl time.Duration) error
	MGet(ctx context.Context, key string, rcv interface{}) error
	Decr(ctx context.Context, key string, val int, ttl time.Duration) (int, error)
	Incr(ctx context.Context, key string, val int, ttl time.Duration) (int, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	Delete(ctx context.Context, key string) error
}

var impl Interface

func Register(i Interface) {
	if impl != nil {
		panic("multiple call")
	}
	impl = i
}

func BitSets(ctx context.Context, key string, val bool, ttl time.Duration, idx ...int) error {
	return trace(ctx, key)(impl.BitSets(ctx, key, val, ttl, idx...))

}

func BitSet(ctx context.Context, key string, idx int, val bool, ttl time.Duration) error {
	return trace(ctx, key)(impl.BitSet(ctx, key, idx, val, ttl))
}

func BitGet(ctx context.Context, key string, idx int) (bool, error) {
	f := trace(ctx, key)
	b, err := impl.BitGet(ctx, key, idx)
	return b, f(err)
}

func Set(ctx context.Context, key string, val string, ttl time.Duration) error {
	return trace(ctx, key)(impl.Set(ctx, key, val, ttl))
}

func Get(ctx context.Context, key string, rcv *string) error {
	return trace(ctx, key)(impl.Get(ctx, key, rcv))
}

func MSet(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	return trace(ctx, key)(impl.MSet(ctx, key, val, ttl))
}

func MGet(ctx context.Context, key string, rcv interface{}) error {
	return trace(ctx, key)(impl.MGet(ctx, key, rcv))
}

func Decr(ctx context.Context, key string, val int, ttl time.Duration) (int, error) {
	f := trace(ctx, key)
	t, err := impl.Incr(ctx, key, -val, ttl)
	return t, f(err)
}

func Incr(ctx context.Context, key string, val int, ttl time.Duration) (int, error) {
	f := trace(ctx, key)
	t, err := impl.Incr(ctx, key, val, ttl)
	return t, f(err)
}

func TTL(ctx context.Context, key string) (time.Duration, error) {
	f := trace(ctx, key)
	t, err := impl.TTL(ctx, key)
	return t, f(err)
}

func Delete(ctx context.Context, key string) error {
	return trace(ctx, key)(impl.Delete(ctx, key))
}

func trace(ctx context.Context, key string) func(err error) error {
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		return func(err error) error {
			return err
		}
	}
	ch := opentracing.StartSpan(impl.Name(), opentracing.ChildOf(sp.Context()))
	logs := []log.Field{log.String(impl.Name()+"_key", key)}
	return func(e error) error {
		if e != nil {
			logs = append(logs, log.Error(e))
			ch.SetTag("error", true)
		}
		ch.LogFields(logs...)
		ch.Finish()
		return e
	}
}
