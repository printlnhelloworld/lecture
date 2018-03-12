package middlewares

import (
	"net/http"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	//NameLecture 讲座名字
	NameLecture = "Lecture"
)

//LectureMustBeExist 讲座必须存在
func LectureMustBeExist(name string) func(*gin.Context) {
	return func(c *gin.Context) {
		lecid := GetLectureIDFromContext(c, name)

		lec, err := model.GetLectureByID(lecid)
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "NotFoundErr",
				"msg":    "讲座 " + c.Param(name) + " 不存在",
			})
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": "DatabaseErr",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
			return
		}
		c.Set(NameLecture, lec)
	}
}

//MustBeLectureOwner 必须是讲座创建者
func MustBeLectureOwner(name string) func(*gin.Context) {
	return func(c *gin.Context) {
		if lecif, exist := c.Get(NameLecture); exist {
			lec := lecif.(*model.Lecture)
			if lec.UserID == GetUserIDFromContext(c) {
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status": "Forbieen",
			"msg":    "你不是讲座所有者，拒绝操作",
		})
	}
}

//GetLectureIDFromContext 从上下文中获取讲座ID
func GetLectureIDFromContext(c *gin.Context, name string) int {
	lectureidStr, _ := c.Get(name)
	return lectureidStr.(int)
}

//GetLectureFromContext 从上下文中获取讲座信息
func GetLectureFromContext(c *gin.Context) *model.Lecture {
	lecif, _ := c.Get(NameLecture)
	return lecif.(*model.Lecture)
}

//LectureMustBeNotFinished 讲座必须未结束
func LectureMustBeNotFinished() func(*gin.Context) {
	return func(c *gin.Context) {
		lec := GetLectureFromContext(c)
		if lec.Finished {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": "Forbidden",
				"msg":    "讲座已经结束，禁止操作",
			})
		}
	}
}
