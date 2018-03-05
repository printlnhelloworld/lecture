package routers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

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
		if limit > 50 {
			limit = 50
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
				Status:  getLectureStatus(lec, now),
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

//PatchLectureByID 修改讲座，不用带上全部参数
func PatchLectureByID() func(*gin.Context) {
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

	}
}

//GetlectureByID 获取特定讲座
func GetlectureByID() func(*gin.Context) {
	return func(c *gin.Context) {
		lectureidStr, _ := c.Get("lectureid")
		lectureid := lectureidStr.(int)

		lec, _ := model.GetLectureByID(lectureid)

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
				"status":        getLectureStatus(*lec, time.Now()),
				"createAt":      lec.CreateAt.Unix(),
				"finished":      lec.Finished,
				"finishedAt":    lec.FinishedAt.Unix(),
				"canSignin":     false, //TODO 未实现
				"remark":        lec.Remark,
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

//lectureCode 讲座的状态
var lectureCodeMap = map[int]lectureCode{}

type lectureCode struct {
	canSign  bool      //是否能签到
	isEnd    bool      //是否结束
	code     string    //签到码
	expireAt time.Time //签到码过期时间
}

//GetLectureCodeByID 生成特定讲座的签到码
func GetLectureCodeByID() func(*gin.Context) {
	return func(c *gin.Context) {
		lid := 0 //TODO
		initLectureCode(lid)
	}
}

func initLectureCode(lid int) {
	if _, ok := lectureCodeMap[lid]; ok {
		return
	}

	l, _ := model.GetLectureByID(lid)
	lectureCodeMap[lid] = lectureCode{
		canSign:  false,
		isEnd:    l.Finished,
		code:     "",
		expireAt: time.Now(),
	}
}

//AddSigninRecordLecturesByID 添加特定讲座签到记录
func AddSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {
		type record struct {
			Type *string `json:"type" binding:"required"`
			Code *string `json:"code"`
			ID   *string `json:"id"`
			Name *string `json:"name"`
		}
		r := record{}
		if err := c.ShouldBindWith(&r, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "参数 type 是必须包含的",
				"err":    err.Error(),
			})
			return
		}
		lectureidStr, _ := c.Get("lectureid")
		lid := lectureidStr.(int)

		lecture, _ := model.GetLectureByID(lid)
		switch *r.Type {
		case "byhand":
			userid, _ := c.Get("UserID")
			if lecture.UserID != userid {
				c.JSON(http.StatusForbidden, gin.H{
					"status": "ok",
					"msg":    "只有讲座创建者才能手动添加签到记录",
				})
			} else {
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
						"data": map[string]interface{}{
							"lecture_id": lr.LectureID,
							"user_id":    lr.UserID,
							"type":       lr.Type,
							"createAt":   lr.CreateAt.Unix(),
							"remark":     lr.Remark,
						},
					})
				}
			}
		case "qcode", "code":
			c.JSON(http.StatusNotImplemented, gin.H{
				"satatus": "ok",
				"msg":     "ok",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "type 值不对，必须是 byhand/qcode/code 中的一种",
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
				"name":     "", //TODO 实现获取名字
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

//GetOneSigninRecordLecturesByID 获取特定讲座的特定用户签到记录
func GetOneSigninRecordLecturesByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		lectureidStr, _ := c.Get("lectureid")
		lid := lectureidStr.(int)
		userid := c.Param("userid")
		lr, err := model.GetLectureRecord(lid, userid)
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
				"data":   lr,
			})
		}
	}
}

func getLectureStatus(lec model.Lecture, now time.Time) string {
	if lec.Finished {
		return "ended"
	} else if now.Before(lec.StartAt) {
		return "prepare"
	}
	return "running"
}
