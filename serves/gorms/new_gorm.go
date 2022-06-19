package gorms

import (
	"database/sql"

	"github.com/wenerme/wego/confs"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGorm(db *sql.DB, conf *confs.DatabaseConf) (gdb *gorm.DB, err error) {
	gc := &gorm.Config{
		Logger:                                   newGLogger(conf),
		DisableForeignKeyConstraintWhenMigrating: conf.GORM.DisableForeignKeyConstraintWhenMigrating,
	}

	switch conf.Type {
	case "sqlite":
		gdb, err = gorm.Open(sqlite.Dialector{
			Conn: db,
		}, gc)
	case "postgres", "postgresql":
		gdb, err = gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), gc)
	}
	return
}
