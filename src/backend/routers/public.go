package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetPublicAgreement 获取用户协议
func GetPublicAgreement(agreement []string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"data":   agreement,
		})
	}
}

//GetLectureTypes 获取讲座类型
func GetLectureTypes() func(*gin.Context) {
	return func(c *gin.Context) { //TODO 实现
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	}
}
