package routers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//GetLectures 获取讲座列表
func GetLectures() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	}
}

//CreateLecture 创建讲座
func CreateLecture() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//PutLectureByID 修改讲座，不用带上全部参数
func PutLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//GetlectureByID 获取特定讲座
func GetlectureByID() func(*gin.Context) {
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

//DeleteLectureByID 删除特定讲座
func DeleteLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//GenerateLectureByID 生成特定讲座的签到码
func GenerateLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//AddSigninRecordLecturesByID 添加特定讲座签到记录
func AddSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//GetSigninRecordLecturesByID 获取特定讲座签到记录
func GetSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//DeleteOneSigninRecordLecturesByID 删除特定讲座签到记录
func DeleteOneSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
