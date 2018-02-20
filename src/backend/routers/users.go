package routers

import "github.com/gin-gonic/gin"

//获取用户列表
func GetUsers() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//获取用户信息
func GetUserByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//获取用户的参加的讲座列表
func GetLecturesByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//获取用户的参加的特定讲座
func GetLectureByLectureIDByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//登录
func AddTokensByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//登录列表
func GetTokensByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//登出
func DeleteTokenByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
