package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouters() *gin.Engine {
	r := gin.Default()
	r.Use(corsHeader)

	apiv1 := r.Group("/api/v1", casAuth("http://cas.hdu.edu.cn/cas/"))
	apiv1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	})

	//讲座
	lectures := apiv1.Group("/lectures")
	lectures.GET("", getLectures())
	lectures.POST("", createLecture())
	lectures.PATCH("/:lectureid", putLectureByID())
	lectures.GET("/:lectureid", getlectureByID())
	lectures.DELETE("/:lectureid", deleteLectureByID())
	lectures.POST("/:lectureid/siginCode", generateLectureByID())
	lectures.POST("/:lectureid/users", addSigninRecordLecturesByID())
	lectures.GET("/:lectureid/users", getSigninRecordLecturesByID())
	lectures.DELETE("/:lectureid/user/:userid", deleteOneSigninRecordLecturesByID())

	//用户
	users := apiv1.Group("/users")
	users.GET("", getUsers())
	users.GET("/:userid", getUserByID())
	users.GET("/:userid/lectures", getLecturesByUserID())
	users.GET("/:userid/lectures/:lectureid", getLectureByLectureIDByUserID())
	users.POST("/:userid/tokens", addTokensByUserID())
	users.GET("/:userid/tokens", getTokensByUserID())
	users.DELETE("/:userid/token/:token", deleteTokenByUserID())

	//管理员
	admin := apiv1.Group("/admin")
	admin.GET("/users", getAdminUsers())
	admin.POST("/users", addAdminUser())
	admin.PATCH("/users/:userid", patchAdminUser())
	admin.DELETE("/users/:userid", deleteAdminUser())
	admin.GET("/output", adminOutput())
	admin.GET("/record", adminRecords())

	//公告
	ann := apiv1.Group("/announcements")
	ann.GET("", getAnnouncements())
	ann.POST("", createAnnouncements())
	ann.GET("/:annoncementid", getAnnouncementByID())
	ann.DELETE("/:annoncementid", deleteAnnouncementByID())
	ann.PUT("/:annoncementid", putAnnouncements())
	return r
}
