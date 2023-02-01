package main

import (
	"github.com/hihibug/microdule/core/etcd"
	"github.com/hihibug/microdule/core/redis"
	"github.com/hihibug/microdule/core/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"log"
	"microdule-layout/internal/config"
	"microdule-layout/internal/global"
	"microdule-layout/pkg"
	"os"
)

var (
	g errgroup.Group
)

func NewService(glb *pkg.Global, redis redis.Redis, db *gorm.DB, etcd etcd.Etcd) *pkg.App {
	return &pkg.App{
		Db:     db,
		Etcd:   etcd,
		Log:    glb.Srv.Options().Log.Client(),
		Srv:    glb.Srv,
		Rpc:    glb.Srv.Rpc().Client(),
		Rest:   glb.Srv.Rest().Client(),
		Redis:  redis,
		Config: glb.Config,
	}
}

func main() {
	//初始化全局
	app, err := InitService()
	if err != nil {
		log.Fatal(err)
	}

	//程序结束释放链接
	defer app.Srv.Close()

	//全局调用声明
	global.Data = app

	//http服务
	g.Go(func() error {
		config.RegisterRoute(app.Rest)
		err := app.Rest.Run()
		if err != nil {
			return err
		}
		return nil
	})

	//rpc服务
	g.Go(func() error {
		ip, _ := utils.ExternalIP()
		app.Config.Rpc.IP = ip
		err := config.RegisterGrpc(app.Rpc.RpcSrv)
		if err != nil {
			return err
		}
		err = app.Rpc.Run()
		if err != nil {
			return err
		}
		return nil
	})

	//监听错误
	if err := g.Wait(); err != nil {
		app.Log.Error("启动失败", zap.Any("err", err))
		os.Exit(1)
	}
}
