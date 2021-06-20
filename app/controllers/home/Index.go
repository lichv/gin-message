package Home

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200,gin.H{
		"state":2000,
		"message":"success",
	})
}
