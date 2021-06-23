package message

import (
	"gin-message/app/services"
	"github.com/gin-gonic/gin"
	lichv "github.com/lichv/go"
	"math/rand"
	"time"
)

func SendVerifyCode(c *gin.Context)  {
	target := c.DefaultQuery("target","")
	if target ==""{
		target = c.DefaultPostForm("target","")
	}
	code := c.DefaultQuery("code","")
	if code == ""{
		code = c.DefaultPostForm("code","")
	}
	if code == ""{
		rand.Seed(time.Now().Unix())
		code = lichv.Strval(rand.Intn(9000)+1000)
	}
	ptime := c.DefaultQuery("ptime","")
	if ptime==""{
		ptime = c.DefaultPostForm("ptime","")
	}

	message, err := services.AddChuanglanMessageWithTemplate(lichv.Strval(target), "3df5erf",ptime, map[string]interface{}{"code":code})
	if err != nil{
		c.JSON(200,gin.H{
			"state":4001,
			"message":err.Error(),
		})
	}

	c.JSON(200,gin.H{
		"state":2000,
		"message":"发送成功",
		"data":message,
	})
}
func SendSmsByChuanglan(c *gin.Context)  {
	target := c.DefaultQuery("target","")
	if target ==""{
		target = c.DefaultPostForm("target","")
	}
	content := c.DefaultQuery("content","")
	if content == ""{
		content = c.DefaultPostForm("code","")
	}

	ptime := c.DefaultQuery("ptime","")
	if ptime==""{
		ptime = c.DefaultPostForm("ptime","")
	}

	message, err := services.AddChuanglanMessage(lichv.Strval(target), content,ptime)
	if err != nil{
		c.JSON(200,gin.H{
			"state":4001,
			"message":err.Error(),
		})
	}

	c.JSON(200,gin.H{
		"state":2000,
		"message":"发送成功",
		"data":message,
	})
}
