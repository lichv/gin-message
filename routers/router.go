package routers

import (
	Home "gin-message/app/controllers/home"
	"gin-message/app/middlewares"
	"gin-message/utils/setting"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"path"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(favicon.New(path.Join(setting.AppSetting.RootPath, "favicon.ico")))
	r.LoadHTMLGlob("./public/*.html")
	r.Static("/static", "./public/static/")
	r.Use(middlewares.Cors())

	r.GET("/", Home.Index)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewares.JWT())
	{
		apiv1.GET("/user", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"state":2000,
				"message":"ok",
			})
		})
	}

	return r
}
