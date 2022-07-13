package redisx

import (
	"github.com/go-redis/redis/v8"
	"github.com/wenerme/wego/contextx"
)

var (
	ClientKey   = contextx.NewKey[*redis.Client](&contextx.KeyOptions{Private: true, Name: "redisx.Client"})
	FromContext = ClientKey.Get
	NewContext  = ClientKey.NewContext
)
