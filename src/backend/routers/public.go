package routers

import (
	"net/http"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
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
func GetLectureTypes(conf *conf.Conf) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"data": func(m map[string]string) []string {
				tmp := []string{}
				for _, s := range m {
					tmp = append(tmp, s)
				}
				return tmp
			}(conf.Unit),
		})
	}
}
