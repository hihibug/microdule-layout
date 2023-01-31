package pkg

import (
	"github.com/google/wire"
	"github.com/hihibug/microdule"
	"github.com/hihibug/microdule/core/etcd"
)

var ProviderETCD = wire.NewSet(InitEtcd)

// InitEtcd 初始化etcd
func InitEtcd(global *Global) etcd.Etcd {
	global.Config.Etcd.Log = global.Srv.Options().Log.Client()

	global.Srv.Init(microdule.Etcd(global.Config.Etcd))
	return global.Srv.Options().Etcd
}
