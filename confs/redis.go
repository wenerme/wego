package confs

//go:generate gomodifytags -file=redis.go -w -all -add-tags yaml -transform snakecase --skip-unexported -add-options yaml=omitempty

import (
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type RedisConf struct {
	DSN          string            `env:"DSN" yaml:"dsn,omitempty"`
	Address      string            `env:"ADDRESS" yaml:"address,omitempty"`
	Username     string            `env:"USERNAME" yaml:"username,omitempty"`
	Password     string            `env:"PASSWORD" yaml:"password,omitempty"`
	Database     int               `env:"DATABASE" yaml:"database,omitempty"`
	DialTimeout  time.Duration     `env:"DIAL_TIMEOUT" yaml:"dial_timeout,omitempty"`
	ReadTimeout  time.Duration     `env:"READ_TIMEOUT" yaml:"read_timeout,omitempty"`
	WriteTimeout time.Duration     `env:"WRITE_TIMEOUT" yaml:"write_timeout,omitempty"`
	Attributes   map[string]string `envPrefix:"ATTR_" yaml:"attributes,omitempty"`
}

//nolint:gocyclo
func (rc *RedisConf) ApplyDNS() {
	if rc.DSN == "" {
		return
	}
	for _, v := range strings.Split(rc.DSN, " ") {
		sp := strings.Split(v, "=")
		if len(sp) != 2 {
			log.Warn().Str("entry", v).Msg("invalid dsn entry")
			continue
		}
		k := strings.TrimSpace(sp[0])
		val := strings.TrimSpace(sp[1])
		switch {
		case k == "address" && rc.Address == "":
			rc.Address = val
		case k == "username" && rc.Username == "":
			rc.Username = val
		case k == "password" && rc.Password == "":
			rc.Password = val
		case k == "database" && rc.Database == 0:
			rc.Database, _ = strconv.Atoi(val)
		case k == "read_timeout" && rc.ReadTimeout == 0:
			rc.ReadTimeout, _ = time.ParseDuration(val)
		case k == "write_timeout" && rc.WriteTimeout == 0:
			rc.WriteTimeout, _ = time.ParseDuration(val)
		case k == "dial_timeout" && rc.DialTimeout == 0:
			rc.DialTimeout, _ = time.ParseDuration(val)
		default:
			log.Warn().Str("entry", v).Msg("unknown dsn entry")
		}
	}
}
