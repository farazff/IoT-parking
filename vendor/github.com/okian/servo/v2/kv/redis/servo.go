package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/mediocregopher/radix/v3"
	"github.com/okian/servo/v2/kv"
	"github.com/spf13/viper"
)

type service struct {
}

func (k *service) Name() string {
	return "redis"
}

func (k *service) Initialize(_ context.Context) error {
	p, err := connection()

	if err != nil {
		return err
	}
	pool = p
	kv.Register(k)
	return pool.Do(radix.Cmd(nil, "PING"))
}

func (k *service) Finalize() error {
	return pool.Close()
}

func (k *service) Healthy(_ context.Context) (interface{}, error) {
	return nil, pool.Do(radix.Cmd(nil, "PING"))
}

func (k *service) Ready(_ context.Context) (interface{}, error) {
	return nil, pool.Do(radix.Cmd(nil, "PING"))
}

type pkgError string

func (p pkgError) Error() string {
	return string(p)
}

const (
	ErrorConnectionField pkgError = "redis connection failed"
	ErrorInvalidHost     pkgError = "redis host is invalid"
	ErrorInvalidPort     pkgError = "redis port is invalid"
)

var (
	pool *radix.Pool
)

const (
	user = "redis_user"
	pass = "redis_pass"
	db   = "redis_db"
	host = "redis_host"
	port = "redis_port"
)

func connection() (*radix.Pool, error) {
	var opt = []radix.DialOpt{
		radix.DialTimeout(time.Second * 10),
	}
	user := viper.GetString(user)
	pass := viper.GetString(pass)
	switch {
	case user != "" && pass != "":
		opt = append(opt, radix.DialAuthUser(user, pass))
	case pass != "":
		opt = append(opt, radix.DialAuthPass(pass))
	}
	if db := viper.GetInt(db); db != 0 {
		opt = append(opt, radix.DialSelectDB(db))
	}

	host := viper.GetString(host)
	port := viper.GetString(port)
	addr := fmt.Sprintf("%s:%s", host, port)
	var connfunc = func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr, opt...)
	}
	return radix.NewPool("tcp", addr, 20, radix.PoolConnFunc(connfunc))
}
