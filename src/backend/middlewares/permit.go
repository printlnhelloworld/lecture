package middlewares

import (
	"net/http"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
	"github.com/gin-gonic/gin"
)

const (
	//PermitSiteAdmin 站点管理权
	PermitSiteAdmin = "siteAdmin"
	//PermitLectureCreate 讲座创建权
	PermitLectureCreate = "lectureCreate"
	//PermitLectureAgree 讲座同意办理权
	PermitLectureAgree = "lectureAgree"
	//PermitRecordView 所有签到记录浏览权
	PermitRecordView = "RecordView"
)

var unitMap map[string]string

//LoadPermits 加载用户权限
func LoadPermits() func(*gin.Context) {
	return func(c *gin.Context) {
		if u, exist := c.Get("UserID"); exist {
			ps := model.GetUserPermits(u.(string))
			userinfo, _ := model.GetUserByID(u.(string))
			if isTeacher(userinfo.Type, userinfo.UnitID) {
				if !havePermit(PermitLectureCreate, ps) {
					ps = append(ps, PermitLectureCreate)
				}
			}
			c.Set("Permits", ps)
		}
	}
}

func permitToMsg(permit string) (msg string) {
	switch permit {
	case PermitLectureAgree:
		msg = "讲座同意办理权"
	case PermitLectureCreate:
		msg = "讲座创建权"
	case PermitRecordView:
		msg = "签到记录浏览权"
	case PermitSiteAdmin:
		msg = "站点管理权"
	}
	return
}

//RequirePermitOr 需要权限，权限之间或关系，只要满足一个就可以了。
func RequirePermitOr(rps ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		if len(rps) == 0 {
			return
		}
		if hpif, exist := c.Get("Permits"); exist {
			if hpif != nil {
				hps := hpif.([]string)
				for _, rp := range rps {
					if havePermit(rp, hps) {
						return
					}
				}
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"status": "Forbidden",
					"msg":    "你需要 " + permitsToMsg(rps) + " 中的一种",
				})
			} else {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"status": "Forbidden",
					"msg":    "你需要 " + permitsToMsg(rps) + " 中的一种",
				})
			}
		}

	}
}

//RequirePermitAnd 需要权限，权限之间与关系，只要满足所有。
func RequirePermitAnd(rps ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		if len(rps) == 0 {
			return
		}
		if hpif, exist := c.Get("Permits"); exist {
			if hpif != nil {
				hps := hpif.([]string)
				matchAll := true
				for _, rp := range rps {
					if !havePermit(rp, hps) {
						matchAll = false
						break
					}
				}
				if matchAll {
					return
				}
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"status": "Forbidden",
					"msg":    "你需要 " + permitsToMsg(rps) + " 中所有权限",
				})
			} else {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"status": "Forbidden",
					"msg":    "你需要 " + permitsToMsg(rps) + " 中所有权限",
				})
			}
		}
	}
}

//havePermit 判断是否有特定权限
func havePermit(p string, ps []string) bool {
	for _, tmp := range ps {
		if p == tmp {
			return true
		}
	}
	return false
}

//permitsToMsg 权限提醒
func permitsToMsg(ps []string) (msg string) {
	for _, tmp := range ps {
		msg += permitToMsg(tmp) + " "
	}
	return
}

//PermitsToMap 权限数组变 map
func PermitsToMap(ps []string) gin.H {
	return gin.H{
		PermitSiteAdmin:     havePermit(PermitSiteAdmin, ps),
		PermitLectureCreate: havePermit(PermitLectureCreate, ps),
		PermitLectureAgree:  havePermit(PermitLectureAgree, ps),
		PermitRecordView:    havePermit(PermitRecordView, ps),
	}
}

func isTeacher(usertype, unitid string) bool {
	if usertype != "3" {
		return false
	}
	if _, ok := unitMap[unitid]; ok {
		return true
	}
	return false
}

//SetUnitMap 设置学院列表
func SetUnitMap(m map[string]string) {
	unitMap = m
}
