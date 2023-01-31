package pkg

import (
	"github.com/google/wire"
	"github.com/hihibug/microdule"
	"github.com/hihibug/microdule/core/gorm"
	"github.com/hihibug/microdule/core/zap"
	gormClient "gorm.io/gorm"
)

var ProviderDb = wire.NewSet(InitOrm)

func InitOrm(global *Global) *gormClient.DB {
	gormConf := gorm.GetGormConfigStruct()
	//对log.New函数的再次封装，从而实现是否通过zap打印日志
	gorm.LogGorm(
		global.Config.DB.LogMode,
		gormConf,
		gorm.SetGormLogZap(zap.NewZapWriter(global.Srv.Options().Log.Client())),
	)
	//注册gorm
	global.Srv.Init(microdule.Gorm(global.Srv.Options().Config.ConfigToGormMysql(gorm.SetGormConfig(gormConf))))

	return global.Srv.Options().Gorm.Client()
}
