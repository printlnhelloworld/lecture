package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//获取讲座列表
func getLectures() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	}
}

//创建讲座
func createLecture() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//修改讲座，不用带上全部参数
func putLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//获取特定讲座
func getlectureByID() func(*gin.Context) {
	return func(c *gin.Context) {
		time.Sleep(time.Millisecond * 100)
		lectureid, err := strconv.Atoi(c.Param("lectureid"))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status": "err",
				"msg":    "讲座id必须是数字",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":    "ok",
				"msg":       "ok",
				"lectrueid": lectureid,
			})
		}
	}
}

//删除特定讲座
func deleteLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//生成特定讲座的签到码
func generateLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//添加特定讲座签到记录
func addSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//获取特定讲座签到记录
func getSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//删除特定讲座签到记录
func deleteOneSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
