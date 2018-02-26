package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetPublicAgreement 获取公共信息
func GetPublicAgreement(agreement []string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"data":   agreement,
		})
	}
}
