package confs

import "time"

type DatabaseConf struct {
	Type         string   `env:"TYPE" yaml:"type,omitempty"`
	Driver       string   `env:"DRIVER" yaml:"driver,omitempty"`
	Database     string   `env:"DATABASE" yaml:"database,omitempty"`
	Username     string   `env:"USERNAME" yaml:"username,omitempty"`
	Password     string   `env:"PASSWORD" yaml:"password,omitempty"`
	Host         string   `env:"HOST" yaml:"host,omitempty"`
	Port         int      `env:"PORT" yaml:"port,omitempty"`
	Schema       string   `env:"SCHEMA" yaml:"schema,omitempty"`
	CreateSchema bool     `env:"CREATE_SCHEMA" yaml:"create_schema,omitempty"`
	DSN          string   `env:"DSN" yaml:"dsn,omitempty"`
	GORM         GORMConf `envPrefix:"GORM_" yaml:"gorm,omitempty"`

	DriverOptions DatabaseDriverOptions `envPrefix:"DRIVER_" yaml:"driver_options,omitempty"`
	Attributes    map[string]string     `envPrefix:"ATTR_" yaml:"attributes,omitempty"` // ConnectionAttributes
}

type DatabaseDriverOptions struct {
	MaxIdleConnections int            `env:"MAX_IDLE_CONNS" yaml:"max_idle_connections,omitempty"`
	MaxOpenConnections int            `env:"MAX_OPEN_CONNS" yaml:"max_open_connections,omitempty"`
	ConnMaxIdleTime    time.Duration  `env:"MAX_IDLE_TIME" yaml:"conn_max_idle_time,omitempty"`
	ConnMaxLifetime    *time.Duration `env:"MAX_LIVE_TIME" yaml:"conn_max_lifetime,omitempty"`
}

type SQLLogConf struct {
	SlowThreshold  time.Duration `env:"SLOW_THRESHOLD" yaml:"slow_threshold,omitempty"`
	IgnoreNotFound bool          `env:"IGNORE_NOT_FOUND" yaml:"ignore_not_found,omitempty"`
	Debug          bool          `env:"DEBUG" yaml:"debug,omitempty"`
}

type GORMConf struct {
	DisableForeignKeyConstraintWhenMigrating bool       `yaml:"disable_foreign_key_constraint_when_migrating,omitempty"`
	Log                                      SQLLogConf `envPrefix:"LOG_" yaml:"log,omitempty"`
}
