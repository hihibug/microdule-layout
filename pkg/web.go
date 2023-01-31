package pkg

import (
	"github.com/google/wire"
	"github.com/hihibug/microdule/web"
)

var ProviderWeb = wire.NewSet(InitRest)

func InitRest(global *Global) *web.Gin {
	rest := global.Srv.Rest().Client()
	InitRoute(rest)
	return rest
}
