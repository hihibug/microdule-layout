package main

import (
	"github.com/hihibug/microdule/core/etcd"
	"github.com/hihibug/microdule/core/redis"
	"github.com/hihibug/microdule/web"
	"gorm.io/gorm"
	"log"
	"microdule-layout/pkg"
)

func NewService(glb *pkg.Global, redis redis.Redis, db *gorm.DB, etcd etcd.Etcd, rest *web.Gin) *pkg.App {
	return &pkg.App{
		Db:    db,
		Etcd:  etcd,
		Log:   glb.Srv.Options().Log.Client(),
		Srv:   glb.Srv,
		Rest:  rest,
		Redis: redis,
	}
}

func main() {
	//初始化全局
	app, err := InitService()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Srv.Close()

	pkg.Data = app
	app.Rest.Run()
}
