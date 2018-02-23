package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"git.hduhelp.com/hduhelper/lecture/src/backend/middlewares"
)

//SetupRouters 初始化路由
func SetupRouters() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsHeader)

	apiv1 := r.Group("/api/v1", middleware.CasAuth("http://cas.hdu.edu.cn/cas/"))
	apiv1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	})

	//讲座
	lectures := apiv1.Group("/lectures")
	lectures.GET("", GetLectures())
	lectures.POST("", CreateLecture())
	lectures.PATCH("/:lectureid", PatchLectureByID())
	lectures.GET("/:lectureid", GetlectureByID())
	lectures.DELETE("/:lectureid", DeleteLectureByID())
	lectures.POST("/:lectureid/siginCode", GenerateLectureByID())
	lectures.POST("/:lectureid/users", AddSigninRecordLecturesByID())
	lectures.GET("/:lectureid/users", GetSigninRecordLecturesByID())
	lectures.DELETE("/:lectureid/user/:userid", DeleteOneSigninRecordLecturesByID())

	//用户
	users := apiv1.Group("/users")
	users.GET("", GetUsers())
	users.GET("/:userid", GetUserByID())
	users.GET("/:userid/lectures", GetLecturesByUserID())
	users.GET("/:userid/lectures/:lectureid", GetLectureByLectureIDByUserID())
	users.POST("/:userid/tokens", AddTokensByUserID())
	users.GET("/:userid/tokens", GetTokensByUserID())
	users.DELETE("/:userid/token/:token", DeleteTokenByUserID())

	//管理员
	admin := apiv1.Group("/admin")
	admin.GET("/users", GetAdminUsers())
	admin.POST("/users", AddAdminUser())
	admin.PATCH("/users/:userid", PatchAdminUser())
	admin.DELETE("/users/:userid", DeleteAdminUser())
	admin.GET("/output", AdminOutput())
	admin.GET("/record", AdminRecords())

	//公告
	ann := apiv1.Group("/announcements")
	ann.GET("", GetAnnouncements())
	ann.POST("", CreateAnnouncements())
	ann.GET("/:announcementid", GetAnnouncementByID())
	ann.DELETE("/:announcementid", DeleteAnnouncementByID())
	ann.PUT("/:announcementid", PutAnnouncementByID())
	return r
}
