package pkg

import (
	"github.com/google/wire"
	"github.com/hihibug/microdule"
	"github.com/hihibug/microdule/core/etcd"
	"github.com/hihibug/microdule/core/redis"
	"github.com/hihibug/microdule/core/viper"
	"github.com/hihibug/microdule/rpc"
	"github.com/hihibug/microdule/web"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ProviderGlobal = wire.NewSet(InitGlobal)
)

type App struct {
	Srv    microdule.Service
	Db     *gorm.DB
	Etcd   etcd.Etcd
	Log    *zap.Logger
	Rest   *web.Gin
	Redis  redis.Redis
	Rpc    *rpc.Grpc
	Config *viper.Config
}

type Global struct {
	Srv    microdule.Service
	Config *viper.Config
}

// InitGlobal 初始化全局
func InitGlobal() (*Global, error) {
	s := microdule.NewService(
		microdule.Config("./config.yml"),
		microdule.Name("microdule-layout"),
	)

	return &Global{
		Srv:    s,
		Config: s.Options().Config.Data,
	}, nil
}
