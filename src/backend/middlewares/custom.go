package middlewares

import (
	"net/http"
	"strconv"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"

	"github.com/gin-gonic/gin"
)

//CorsHeader 添加响应头，实现跨域使用、同时实现允许使用 REST 中的方法
func CorsHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Authorization,Content-type")
	c.Header("X-Powered-By", "php")
	return
}

//Auth 认证授权等
func Auth(prefix string, unAuthPath ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		for _, uap := range unAuthPath {
			if prefix+uap == path {
				return
			}
		}

		if c.Request.Method == "OPTIONS" {
			return
		}

		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
				"msg":    "需要登录",
			})
		} else {
			c.Set("Token", token)
			userid, err := model.GetUserIDByToken(token)
			if err == nil {
				c.Set("UserID", userid)
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status": "badTokenErr",
					"msg":    "错误或过期的token",
					"token":  token,
				})
			}
		}
	}
}

//PathParamMustBeInt 指定 Path 参数必须为数字
func PathParamMustBeInt(args ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		for _, arg := range args {
			tmp, err := strconv.Atoi(c.Param(arg))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"status": "PathParamErr",
					"msg":    arg + "必须是数字",
				})
				return
			}
			c.Set(arg, tmp)
		}
	}
}
