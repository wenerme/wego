package confs

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func (conf *DatabaseConf) BuildDSN() string {
	if conf.DSN != "" {
		return conf.DSN
	}
	driver := conf.Driver

	switch conf.Type {
	case "sqlite":
		if driver == "" {
			driver = "sqlite"
		}
		conf.DSN = buildSqliteDSN(conf)
	case "postgres", "postgresql":
		if driver == "" {
			driver = "pgx"
		}
		conf.DSN = buildPgDSN(conf)
	}

	conf.Driver = driver
	return conf.DSN
}

func buildSqliteDSN(conf *DatabaseConf) (dsn string) {
	dbFile := os.ExpandEnv(conf.Database)
	u := url.URL{
		Scheme: "file",
		Path:   dbFile,
	}
	for k, v := range conf.Attributes {
		u.Query().Add(k, v)
	}
	dsn = u.String()
	return
}

func buildPgDSN(conf *DatabaseConf) (dsn string) {
	o := make(map[string]string)
	o["user"] = conf.Username
	o["password"] = conf.Password
	o["host"] = conf.Host
	if conf.Port != 0 {
		o["port"] = fmt.Sprint(conf.Port)
	}
	o["dbname"] = conf.Database
	o["search_path"] = conf.Schema
	for k, v := range conf.Attributes {
		o[k] = v
	}
	sb := strings.Builder{}
	for k, v := range o {
		if v == "" {
			continue
		}
		sb.WriteString(k)
		sb.WriteRune('=')
		sb.WriteString(v)
		sb.WriteRune(' ')
	}
	dsn = sb.String()
	dsn = regexp.MustCompile(`\s+`).ReplaceAllString(dsn, " ")
	dsn = strings.TrimSpace(dsn)
	return
}
