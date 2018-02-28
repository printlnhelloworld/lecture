package middlewares

import (
	"net/http"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//AnnouncementMustBeExist 公告必须存在
func AnnouncementMustBeExist(name string) func(*gin.Context) {
	return func(c *gin.Context) {
		astr, _ := c.Get(name)
		aid := astr.(int)

		_, err := model.GetAnnouncementByID(aid)
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "NotFoundErr",
				"msg":    "公告 " + c.Param(name) + " 不存在",
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
