package redisx

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/wenerme/wego/confs"
)

// BuildOptions build redis.Options from confs.RedisConf
func BuildOptions(rc *confs.RedisConf) (*redis.Options, error) {
	ro := &redis.Options{
		Addr:         rc.Address,
		Username:     rc.Username,
		Password:     rc.Password,
		DB:           rc.Database,
		DialTimeout:  rc.DialTimeout,
		ReadTimeout:  rc.ReadTimeout,
		WriteTimeout: rc.WriteTimeout,
	}
	return ro, nil
}
