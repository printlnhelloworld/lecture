package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
