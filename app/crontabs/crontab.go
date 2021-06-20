package crontabs

import (
	"github.com/robfig/cron"
	"log"
)

func Setup() {
	c := cron.New()
	c.AddFunc("0 0 */2 * * *", func() {
		log.Println("定时爬虫获取样本跟踪数据")
	})

	c.Run()

	return
}
