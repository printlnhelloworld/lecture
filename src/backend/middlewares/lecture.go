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
		lecid := getLectureID(name, c)

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
			if lec.UserID == getUserID(c) {
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status": "Forbieen",
			"msg":    "你不是讲座所有者，拒绝操作",
		})
	}
}

func getLectureID(name string, c *gin.Context) int {
	lectureidStr, _ := c.Get(name)
	return lectureidStr.(int)
}

func getUserID(c *gin.Context) string {
	userif, _ := c.Get("UserID")
	return userif.(string)
}
