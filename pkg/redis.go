package pkg

import (
	"github.com/google/wire"
	"github.com/hihibug/microdule"
	"github.com/hihibug/microdule/core/redis"
)

var ProviderRedis = wire.NewSet(InitRedis)

// InitRedis 初始化redis
func InitRedis(global *Global) redis.Redis {
	global.Srv.Init(microdule.Redis(nil))
	return global.Srv.Options().Redis
}
