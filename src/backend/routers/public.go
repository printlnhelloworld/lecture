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
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"data": []string{
				"校团委讲座",
				"理学院讲座",
				"计算机学院讲座",
			},
		})
	}
}
