package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//corsHeader 添加响应头，实现跨域使用、同时实现允许使用 REST 中的方法
func corsHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
	return
}

func casAuth(path string) func(*gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		c.Set("token", token)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
				"msg":    "需要登录",
			})
		}
		c.Next()
	}
}
