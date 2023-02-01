package api

import (
	"github.com/gin-gonic/gin"
	"microdule-layout/internal/global"
)

func Index(context *gin.Context) {
	global.Data.Log.Info("test")
}
