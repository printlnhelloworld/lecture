package middlewares

import (
	"net/http"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//LectureMustBeExist 讲座必须存在
func LectureMustBeExist(name string) func(*gin.Context) {
	return func(c *gin.Context) {
		lectureidStr, _ := c.Get(name)
		lecid := lectureidStr.(int)

		_, err := model.GetLectureByID(lecid)
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "NotFoundErr",
				"msg":    "讲座 " + c.Param(name) + " 不存在",
			})
		} else if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": "DatabaseErr",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
		}
	}
}
