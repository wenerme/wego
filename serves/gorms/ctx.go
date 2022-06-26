package gorms

import (
	"github.com/wenerme/wego/contextx"
	"gorm.io/gorm"
)

var (
	DBKey       = contextx.NewKey[*gorm.DB](&contextx.KeyOptions{Private: true, Name: "gorms.DB"})
	FromContext = DBKey.Get
	NewContext  = DBKey.NewContext
)
