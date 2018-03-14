package routers

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"git.hduhelp.com/hduhelper/lecture/src/backend/middlewares"
	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
)

//GetLectures 获取讲座列表
func GetLectures() func(*gin.Context) {
	return func(c *gin.Context) {
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
		next, err1 := strconv.Atoi(c.DefaultQuery("next", "0"))
		if err != nil || err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamError",
				"msg":    "参数limit/next必须为数字",
			})
			return
		}

		if limit > 50 {
			limit = 50
		}

		owner := c.DefaultQuery("owner", "")
		status := c.DefaultQuery("status", "all")

		sort := c.DefaultQuery("sort", "id")
		sortMatch := false
		for _, stype := range []string{"id", "startAt"} {
			if stype == sort {
				sortMatch = true
			}
		}
		if sortMatch == false {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamError",
				"msg":    "参数sort必须 id / startAt",
			})
			return
		}
		type lecture struct {
			ID      int    `json:"id"`
			Topic   string `json:"topic"`
			Type    string `json:"type"`
			Status  string `json:"status"`
			StartAt int64  `json:"startAt"`
		}
		now := time.Now()
		lecs, err := model.GetLectures(limit, next, owner, status, sort, now)
		ls := make([]lecture, 0)
		var newNext int64
		if ll := len(*lecs); ll > 0 {
			switch sort {
			case "id":
				newNext = int64((*lecs)[ll-1].ID)
			case "startAt":
				newNext = (*lecs)[ll-1].StartAt.Unix()
			}
		}
		for _, lec := range *lecs {
			ls = append(ls, lecture{
				ID:      lec.ID,
				Topic:   lec.Topic,
				Type:    lec.Type,
				Status:  getLectureStatus(lec),
				StartAt: lec.StartAt.Unix(),
			})
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "DatabaseError",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"msg":    "ok",
				"next":   newNext,
				"data":   ls,
			})
		}
	}
}

//CreateLecture 创建讲座
func CreateLecture() func(*gin.Context) {
	//TODO 实现教师校验
	type lecture struct {
		Topic        string `json:"topic" binding:"required"`
		Location     string `json:"location" binding:"required"`
		Introduction string `json:"introduction" binding:"required"`
		StartAt      int64  `json:"startAt" binding:"required"`
		Host         string `json:"host" binding:"required"`
		Lecturer     string `json:"lecturer" binding:"required"`
		Type         string `json:"type" binding:"required"`
		Reviewed     bool   `json:"reviewed" binding:"required"`
	}
	return func(c *gin.Context) {
		var lec lecture
		if err := c.ShouldBindWith(&lec, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "参数必须带全",
				"err":    err.Error(),
			})
		} else {
			userid, _ := c.Get("UserID")
			if id, err := model.CreateLecture(
				userid.(string),
				lec.Topic,
				lec.Location,
				lec.Introduction,
				lec.Host,
				lec.Lecturer,
				lec.Type,
				lec.Reviewed,
				time.Unix(lec.StartAt, 0),
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "DatabaseError",
					"msg":    "数据库出现错误",
					"err":    err.Error(),
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{
					"status": "ok",
					"msg":    "ok",
					"id":     id,
				})
			}
		}
	}
}

//PutLectureByID 修改讲座，不用带上全部参数
func PutLectureByID() func(*gin.Context) {
	type lecture struct {
		Topic        *string `json:"topic"`
		Location     *string `json:"location"`
		Introduction *string `json:"introduction"`
		StartAt      *int64  `json:"startAt"`
		Host         *string `json:"host"`
		Lecturer     *string `json:"lecturer"`
		Type         *string `json:"type"`
		Reviewed     *bool   `json:"reviewed"`
	}
	return func(c *gin.Context) {
		var lec lecture
		if err := c.ShouldBindWith(&lec, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "至少要有空的 {} ",
				"err":    err.Error(),
			})
			return
		}
		userid, _ := c.Get("UserID")

		lectureidStr, _ := c.Get("lectureid")
		lecid := lectureidStr.(int)

		oldlec, _ := model.GetLectureByID(lecid)

		//TODO 多个人的权限？
		if oldlec.UserID != userid {
			c.JSON(http.StatusForbidden, gin.H{
				"status": "Forbidden",
				"msg":    "禁止修改，只有创建者可以修改",
			})
			return
		}

		//TODO 无聊的过程
		//TODO 处理CanSignin
		var m = map[string]interface{}{}
		{
			if lec.Topic != nil {
				m["Topic"] = *lec.Topic
			}
			if lec.Location != nil {
				m["Location"] = *lec.Location
			}
			if lec.StartAt != nil {
				m["StartAt"] = time.Unix(*lec.StartAt, 0)
			}
			if lec.Introduction != nil {
				m["Introduction"] = *lec.Introduction
			}
			if lec.Host != nil {
				m["Host"] = *lec.Host
			}
			if lec.Lecturer != nil {
				m["Lecturer"] = *lec.Lecturer
			}
			if lec.Type != nil {
				m["Type"] = *lec.Type
			}
			if lec.Reviewed != nil {
				m["Reviewed"] = *lec.Reviewed
			}
		}
		err := model.PatchLecture(lecid, m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "DatabaseErr",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"msg":    "ok",
			})
		}
	}
}

//UpdateLectureStatusByID 更新讲座状态
func UpdateLectureStatusByID() func(*gin.Context) {
	return func(c *gin.Context) {
		type status struct {
			Status string `json:"status" binding:"required"`
		}
		var s status
		if err := c.ShouldBindWith(&s, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "badRequest",
				"msg":    "参数错误",
				"err":    err.Error(),
			})
			return
		}
		switch s.Status {
		case "signing", "notsigning", "ended":
			if lecif, exist := c.Get(middlewares.NameLecture); exist {
				lec := lecif.(*model.Lecture)
				if lec.Finished {
					c.JSON(http.StatusBadRequest, gin.H{
						"status": "badRequest",
						"msg":    "讲座已经结束、不能进行操作",
					})
					return
				}
				if err := model.UpdateLectureStatus(lec.ID, s.Status); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "databaseErr",
						"msg":    "数据库错误",
						"err":    err.Error(),
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status": "ok",
						"msg":    "ok",
					})
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "ServerError",
					"msg":    "服务器错误，lec 不存在",
				})
			}
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "badStatus",
				"msg":    "状态必须是 signing, notsigning, ended 中的一种",
			})
		}
	}
}

//GetlectureByID 获取特定讲座
func GetlectureByID() func(*gin.Context) {
	return func(c *gin.Context) {
		lectureidStr, _ := c.Get("lectureid")
		lectureid := lectureidStr.(int)
		userid := middlewares.GetUserIDFromContext(c)
		lec, _ := model.GetLectureByID(lectureid)
		lr, err2 := model.GetLectureRecord(lectureid, userid)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"data": gin.H{
				"id":            lec.ID,
				"creatorUserID": lec.UserID,
				"topic":         lec.Topic,
				"location":      lec.Location,
				"introduction":  lec.Introduction,
				"startAt":       lec.StartAt.Unix(),
				"host":          lec.Host,
				"lecturer":      lec.Lecturer,
				"type":          lec.Type,
				"reviewed":      lec.Reviewed,
				"status":        getLectureStatus(*lec),
				"createAt":      lec.CreateAt.Unix(),
				"finished":      lec.Finished,
				"finishedAt":    lec.FinishedAt.Unix(),
				"remark":        lec.Remark,
				"signin": gin.H{
					"isSigned": err2 != gorm.ErrRecordNotFound,
					"signedAt": lr.CreateAt.Unix(),
					"type":     lr.Type,
					"remark":   lr.Remark,
				},
			},
		})

	}
}

//DeleteLectureByID 删除特定讲座
func DeleteLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {
		lectureidStr, _ := c.Get("lectureid")
		lectureid := lectureidStr.(int)

		_ = model.DeleteLectureByID(lectureid) //TODO handle error
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
		})

	}
}

//GetLectureCodeByID 生成特定讲座的签到码
func GetLectureCodeByID() func(*gin.Context) {
	return func(c *gin.Context) {
		code, expireAt := getLectureCodeByIDAndUpdateOnExpired(c)
		c.JSON(http.StatusOK, gin.H{
			"status":     "ok",
			"msg":        "ok",
			"signinCode": code,
			"exipreAt":   expireAt.Unix(),
			"exipreIn":   int(expireAt.Sub(time.Now()).Seconds()),
		})
	}
}

func getLectureCodeFromContext(c *gin.Context) (int, string, time.Time) {
	lec := middlewares.GetLectureFromContext(c)
	return lec.ID, lec.SignCode, lec.SignCodeExpireAt
}

func getLectureCodeByIDAndUpdateOnExpired(c *gin.Context) (string, time.Time) {
	lid, code, expireAt := getLectureCodeFromContext(c)
	if !time.Now().Before(expireAt) {
		code = newSignCode()
		expireAt = time.Now().Add(time.Second * 10)
		model.UpdateLectureSignCode(lid, code, expireAt)
	}
	return code, expireAt
}

func newSignCode() string {
	return strconv.Itoa((rand.Intn(999999) + 100000) % 1000000)
}

//AddLectureSigninRecordByhand 添加特定讲座签到记录
func AddLectureSigninRecordByhand() func(*gin.Context) {
	return func(c *gin.Context) {
		type record struct {
			ID *string `json:"id" binding:"required"`
		}
		r := record{}
		if err := c.ShouldBindWith(&r, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "参数 id 是必须包含的",
				"err":    err.Error(),
			})
			return
		}
		lectureidStr, _ := c.Get("lectureid")
		lid := lectureidStr.(int)

		lecture, _ := model.GetLectureByID(lid)
		//TODO 返回系统中没有用户的情况
		lr, err := model.AddLectureRecord("byhand", *r.ID, lecture.ID)
		if err != nil {
			if strings.Contains(err.Error(), "1062") {
				c.JSON(http.StatusOK, gin.H{
					"status": "CreatedErr",
					"msg":    "已经添加过了",
					"err":    err.Error(),
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "DatabaseErr",
					"msg":    "数据库错误",
					"err":    err.Error(),
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"msg":    "ok",
				"data": gin.H{
					"id":   lr.UserID,
					"name": lr.UserInfo.Name, //TODO 获取姓名
				},
			})
		}
	}
}

//AddLectureSigninRecordByCode 添加特定讲座签到记录
func AddLectureSigninRecordByCode() func(*gin.Context) {
	return func(c *gin.Context) {
		type record struct {
			Code *string `json:"code" binding:"required"`
			Type *string `json:"type"`
		}
		r := record{}
		if err := c.ShouldBindWith(&r, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "参数 code 是必须包含的",
				"err":    err.Error(),
			})
			return
		}
		uid := middlewares.GetUserIDFromContext(c)
		lid := middlewares.GetLectureIDFromContext(c, "lectureid")

		_, code, expireAt := getLectureCodeFromContext(c)
		if code == *r.Code && time.Now().Before(expireAt) {
			//TODO 完善文档
			ty := ""
			if r.Type != nil {
				switch *r.Type {
				case "code", "qcode":
					ty = *r.Type
				default:
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"status": "ParamErr",
						"msg":    "type 必须为 code/qcode",
					})
					return
				}

			}
			_, err := model.AddLectureRecord(ty, uid, lid)
			if err != nil {
				if strings.Contains(err.Error(), "1062") {
					c.JSON(http.StatusOK, gin.H{
						"status": "CreatedErr",
						"msg":    "已经添加过了",
						"err":    err.Error(),
					})
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "DatabaseErr",
						"msg":    "数据库错误",
						"err":    err.Error(),
					})
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
					"msg":    "ok",
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "badCode",
				"msg":    "错误的签到码",
			})
		}
	}
}

//GetSigninRecordLecturesByID 获取特定讲座签到记录
func GetSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {
		lectureidStr, _ := c.Get("lectureid")
		lid := lectureidStr.(int)

		total, lrs := model.GetLectureRecordsByLectureID(lid)
		var tmp []map[string]interface{}
		for _, lr := range lrs {
			tmp = append(tmp, map[string]interface{}{
				"userId":   lr.UserID,
				"name":     lr.UserInfo.Name, //TODO 实现获取名字
				"signedAt": lr.CreateAt.Unix(),
				"type":     lr.Type,
				"remark":   lr.Remark,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"total":  total,
			"data":   tmp,
		})
	}
}

//DeleteOneSigninRecordLecturesByID 删除特定讲座特定用户签到记录
func DeleteOneSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {
		lectureidStr, _ := c.Get("lectureid")
		lid := lectureidStr.(int)

		userid := c.Param("userid")
		err := model.DeleteLectureRecord(lid, userid)
		//TODO 设置只能删除手动添加的记录
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "databaseErr",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"msg":    "ok",
			})
		}
	}
}

func getLectureStatus(lec model.Lecture) string {
	if lec.Finished {
		return "ended"
	} else if lec.CanSignin {
		return "signing"
	}
	return "notsigning"
}
