package routers

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
	"git.hduhelp.com/hduhelper/lecture/src/backend/middlewares"
)

const (
	version   = "beta"
	copyright = "杭电助手 © 版权所有"
)

//SetupRouters 初始化路由
func SetupRouters(conf *conf.Conf) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CorsHeader, middlewares.PHP)
	r.Use(gzip.Gzip(gzip.DefaultCompression)) //gzip压缩
	apiv1 := r.Group("/api/v1",
		middlewares.Auth( //不需要登录的接口
			"/api/v1",        //接口前缀
			"/loginCallback", //登录相关
			"/loginURL",      //登录相关
		),
		middlewares.LoadPermits(),
	)
	apiv1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	})

	apiv1.OPTIONS("/*all", func(c *gin.Context) {}) //cors

	//讲座
	lectures := apiv1.Group("/lectures")
	{
		//TODO 讲座结束后不能进行写操作
		lectureidstr := "lectureid"
		lectures.GET("", GetLectures())
		lectures.POST("",
			middlewares.RequirePermitOr(middlewares.PermitLectureCreate), //必须有创建权限
			CreateLecture(),
		)

		lectures = lectures.Group("",
			middlewares.PathParamMustBeInt(lectureidstr), //讲座id必须为数字
			middlewares.LectureMustBeExist(lectureidstr), //讲座必须存在
		)
		lectures.GET("/:"+lectureidstr, GetlectureByID())
		lectures.GET("/:"+lectureidstr+"/users/:userid", GetOneSigninRecordLecturesByID())
		lectures.PUT("/:"+lectureidstr,
			middlewares.MustBeLectureOwner(lectureidstr), //必须是讲座所有者
			middlewares.LectureMustBeNotFinished(),
			PutLectureByID(),
		) //修改讲座
		lectures.PUT("/:"+lectureidstr+"/status",
			middlewares.MustBeLectureOwner(lectureidstr),
			UpdateLectureStatusByID(),
		)
		lectures.DELETE("/:"+lectureidstr,
			middlewares.MustBeLectureOwner(lectureidstr),
			DeleteLectureByID(),
		)
		lectures.GET("/:"+lectureidstr+"/signinCode",
			middlewares.MustBeLectureOwner(lectureidstr),
			GetLectureCodeByID(),
		) //获取签到码
		lectures.POST("/:"+lectureidstr+"/users/code", AddLectureSigninRecordByCode()) //签到码签到
		lectures.POST("/:"+lectureidstr+"/users/byhand",
			middlewares.MustBeLectureOwner(lectureidstr),
			AddLectureSigninRecordByhand(), //手动签到
		)
		lectures.GET("/:"+lectureidstr+"/users", //特定讲座签到记录
			middlewares.MustBeLectureOwner(lectureidstr),
			GetSigninRecordLecturesByID(),
		)
		lectures.DELETE("/:"+lectureidstr+"/users/:userid",
			middlewares.MustBeLectureOwner(lectureidstr),
			DeleteOneSigninRecordLecturesByID(),
		)
	}
	//用户
	users := apiv1.Group("/user") //这里面的权限不需要特殊处理，因为权限是只是限制到本用户，而本用户是根据token得到的。
	{
		users.GET("/userinfo", GetUserInfo())
		users.POST("/agree", UpdateUserAgree())
		users.GET("/lectures", GetUserLectures())
		users.GET("/tokens", GetUserTokens())
		users.DELETE("/tokens", DeleteUserToken("all"))
		users.DELETE("/tokens/self", DeleteUserToken("self"))
		users.PUT("/tokens/self", UpdateUserTokenRemark())
		users.DELETE("/tokens/other", DeleteUserToken("other"))
	}
	//登录
	apiv1.GET("/loginCallback", UserLoginCallBack(conf))
	apiv1.GET("/loginURL", GetLoginURL(conf))

	//管理员
	admin := apiv1.Group("/admin")
	{
		//TODO 完善接口 特别是权限的授予与收回
		admin.GET("/users", GetAdminUsers())
		admin.POST(
			"/users",
			middlewares.RequirePermitOr(middlewares.PermitSiteAdmin),
			AddAdminUser(),
		)
		admin.PATCH(
			"/users/:userid",
			middlewares.RequirePermitOr(middlewares.PermitSiteAdmin),
			PatchAdminUser(),
		)
		admin.DELETE("/users/:userid",
			middlewares.RequirePermitOr(middlewares.PermitSiteAdmin),
			DeleteAdminUser(),
		)
		admin.GET("/output",
			middlewares.RequirePermitOr(middlewares.PermitRecordView),
			AdminOutput(),
		)
		admin.GET("/record",
			middlewares.RequirePermitOr(middlewares.PermitRecordView),
			AdminRecords(),
		)

	}
	//公告
	ann := apiv1.Group("/announcements")
	{
		anidstr := "announcementid"
		ann.GET("", GetAnnouncements())
		ann.POST("", middlewares.RequirePermitOr(middlewares.PermitSiteAdmin), CreateAnnouncements())
		ann = ann.Group("",
			middlewares.PathParamMustBeInt(anidstr),
			middlewares.AnnouncementMustBeExist(anidstr),
		)
		ann.GET("/:"+anidstr, GetAnnouncementByID()) //TODO remove
		ann.DELETE("/:"+anidstr, middlewares.RequirePermitOr(middlewares.PermitSiteAdmin), DeleteAnnouncementByID())
		ann.PUT("/:"+anidstr, middlewares.RequirePermitOr(middlewares.PermitSiteAdmin), PutAnnouncementByID())
	}
	//公开信息

	public := apiv1.Group("/public")
	public.GET("/agreement", GetPublicAgreement(conf.Agreement))
	public.GET("/system_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"data": gin.H{
				"version":   version,
				"copyright": copyright,
			},
		})
	})
	public.GET("/lecture_type", GetLectureTypes(conf))

	//前端static页面
	r.StaticFS("/static", packr.NewBox("../../../dist/static"))

	front := packr.NewBox("../../../dist")
	handleIndex := func(c *gin.Context) {
		c.Status(http.StatusOK)
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.Write(front.Bytes("index.html"))
	}
	r.GET("index.html", handleIndex)
	r.NoRoute(handleIndex)
	return r
}
