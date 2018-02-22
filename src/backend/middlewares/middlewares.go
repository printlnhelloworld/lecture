package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CorsHeader 添加响应头，实现跨域使用、同时实现允许使用 REST 中的方法
func CorsHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
	return
}

//CasAuth
func CasAuth(path string) func(*gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
				"msg":    "需要登录",
			})
		} else {
			c.Set("token", token)
			//TODO 实现真正的通过token获取学号
			if token == "x" {
				c.Set("UserID", "15051237")
			}
		}
	}
}
