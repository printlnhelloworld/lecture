package main

import "github.com/gin-gonic/gin"

//获取用户列表
func getUsers() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//获取用户信息
func getUserByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//获取用户的参加的讲座列表
func getLecturesByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//获取用户的参加的特定讲座
func getLectureByLectureIDByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//登录
func addTokensByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//登录列表
func getTokensByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//登出
func deleteTokenByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
