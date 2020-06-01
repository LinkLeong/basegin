package router

import (
	"../pak/setting"
	"../router/view"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/statics", "./statics")
	r.StaticFile("/favicon.ico", "./statics/favicon.ico")
	gin.SetMode(setting.ServeSetting.RunMode)
	r.GET("/ddd.html", view.Index)
	r.GET("/:c/index.html", view.Index)
	r.GET("/:c/bbb.html", view.Index)
	return r
}
