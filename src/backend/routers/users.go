package routers

import (
	"net/http"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
	"github.com/gin-gonic/gin"
)

//GetUsers 获取用户列表
func GetUsers() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//GetUserByID 获取用户信息
func GetUserByID() func(*gin.Context) {
	return func(c *gin.Context) {
		userid := c.Param("userid")
		u, err := model.GetUserByID(userid)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"status": "DatabaseError",
				"msg":    "数据库错误",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"stauts": "ok",
				"msg":    "ok",
				"id":     u.UserID,
				"name":   u.Name,
				"agree":  u.Agreed,
			})
		}
	}
}

//GetLecturesByUserID 获取用户的参加的讲座列表
func GetLecturesByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//GetLectureByLectureIDByUserID 获取用户的参加的特定讲座
func GetLectureByLectureIDByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//AddTokensByUserID 登录
func AddTokensByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//GetTokensByUserID 登录列表
func GetTokensByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//DeleteTokenByUserID 登出
func DeleteTokenByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
