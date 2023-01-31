package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hihibug/microdule/web"
)

func InitRoute(rest *web.Gin) {
	a := rest.Route.Group("")
	{
		a.GET("/test", func(context *gin.Context) {
			fmt.Println("test")
			Data.Log.Info("test")
		})
		a.GET("/err", func(c *gin.Context) {
			panic("test")
		})
	}
}
