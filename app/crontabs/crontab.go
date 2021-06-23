package crontabs

import (
	"gin-message/app/services"
	"github.com/robfig/cron"
	"log"
)

func Setup() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("每秒查询一次待发送的信息")
		services.HandleMesssage()
	})

	c.Run()

	return
}
