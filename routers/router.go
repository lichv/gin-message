package routers

import (
	Home "gin-message/app/controllers/home"
	"gin-message/app/controllers/message"
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

	r.Use(middlewares.Cors())

	r.GET("/", Home.Index)

	apiv1 := r.Group("/api/message/v1")
	apiv1.Use(middlewares.Log())
	{
		apiv1.GET("/sms/sendVerifyCode",message.SendVerifyCode)
		apiv1.GET("/sms/sendChuanglan",message.SendSmsByChuanglan)
	}

	return r
}
