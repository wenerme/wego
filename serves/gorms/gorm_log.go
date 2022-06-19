package gorms

import (
	"log"
	"os"
	"time"

	"github.com/wenerme/wego/confs"
	"gorm.io/gorm/logger"
)

func newGLogger(conf *confs.DatabaseConf) logger.Interface {
	return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), newGLoggerConf(conf))
}

func newGLoggerConf(conf *confs.DatabaseConf) logger.Config {
	lc := logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: conf.GORM.Log.IgnoreNotFound,
		Colorful:                  true,
	}
	if conf.Type == "sqlite" && conf.GORM.Log.SlowThreshold == 0 {
		lc.SlowThreshold = time.Second
	} else {
		lc.SlowThreshold = conf.GORM.Log.SlowThreshold
	}
	return lc
}
