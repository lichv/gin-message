package message

import (
	"fmt"
	"gin-message/app/services"
	"gin-message/utils"
	"github.com/gin-gonic/gin"
	lichv "github.com/lichv/go"
)

func sendVerifyCode(c *gin.Context)  {
	input,_ := utils.GetMapFromContext(c)
	target,ok := input["target"]
	if !ok {
		c.JSON(200,gin.H{
			"state":"301",
			"message":"缺失参数",
		})
		return
	}

	fmt.Println(input)

	services.SendSms("chuanglan",lichv.Strval(target),"",input)
	c.JSON(200,gin.H{
		"state":2000,
		"message":"发送成功",
	})
}
