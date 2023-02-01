package config

import (
	"github.com/gin-gonic/gin"
	"github.com/hihibug/microdule/web"
	"microdule-layout/internal/api"
)

func RegisterRoute(rest *web.Gin) {
	a := rest.Route.Group("")
	{
		a.GET("/test", api.Index)
		a.GET("/err", func(c *gin.Context) {
			panic("test")
		})
	}
}
