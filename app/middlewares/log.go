package middlewares

import (
	"gin-message/app/services"
	"gin-message/utils"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		input,_ := utils.GetMapFromContext(c)
		services.LogFromContext(c, path, "", input)
	}
}
