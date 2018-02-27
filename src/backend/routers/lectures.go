package routers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"

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
			if userid, exist := c.Get("UserID"); !exist {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "ServerError",
					"msg":    "服务出现错误",
					"err":    "Lost UserID",
				})
			} else if id, err := model.CreateLecture(
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
		CanSignin    *bool   `json:"canSignin"`
		Finished     *bool   `json:"finished"`
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
		userid, exist := c.Get("UserID")
		if !exist {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "ServerError",
				"msg":    "服务出现错误",
				"err":    "Lost UserID",
			})
			return
		}
		lecid, err := strconv.Atoi(c.Param("lectureid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "参数 lecid 必须是数字",
			})
			return
		}
		oldlec, err := model.GetLectureByID(lecid)
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "NotFoundErr",
				"msg":    "没有数据",
			})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "DatabaseErr",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
			return
		}

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
			if lec.Introduction != nil {
				m["StartAt"] = *lec.StartAt
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
			if lec.Finished != nil {
				m["Finished"] = *lec.Finished
			}
		}
		err = model.PatchLecture(lecid, m)
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

//GetlectureByID 获取特定讲座
func GetlectureByID() func(*gin.Context) {
	return func(c *gin.Context) {
		lectureid, err := strconv.Atoi(c.Param("lectureid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "err",
				"msg":    "讲座id必须是数字",
			})
		} else {
			lec, err := model.GetLectureByID(lectureid)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					c.JSON(http.StatusNotFound, gin.H{
						"status": "NotFoundErr",
						"msg":    "没有数据",
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
						"id":            lec.ID,
						"creatorUserID": lec.UserID,
						"topic":         lec.Topic,
						"location":      lec.Location,
						"startAt":       lec.StartAt.Unix(),
						"host":          lec.Host,
						"lecturer":      lec.Lecturer,
						"type":          lec.Type,
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
	}
}

//DeleteLectureByID 删除特定讲座
func DeleteLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {
		lectureid, err := strconv.Atoi(c.Param("lectureid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "err",
				"msg":    "讲座id必须是数字",
			})
		} else {
			err = model.DeleteLectureByID(lectureid)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					c.JSON(http.StatusNotFound, gin.H{
						"status": "NotFoundErr",
						"msg":    "没有数据",
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
		}
	}
}

//GenerateLectureByID 生成特定讲座的签到码
func GenerateLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {

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
		lid, err := strconv.Atoi(c.Param("lectureid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "讲座id必须是数字",
			})
			return
		}

		lecture, err := model.GetLectureByID(lid)
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "讲座不存在，id 为：" + strconv.Itoa(lid),
			})
			return
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "DatabaseErr",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
			return
		}
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
		lid, err := strconv.Atoi(c.Param("lectureid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "参数 lectureid 必须是数字",
			})
			return
		}
		_, err = model.GetLectureByID(lid)
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "讲座不存在，id 为：" + strconv.Itoa(lid),
			})
			return
		}
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
		lid, err := strconv.Atoi(c.Param("lectureid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "讲座id必须是数字",
			})
			return
		}
		userid := c.Param("userid")
		err = model.DeleteLectureRecord(lid, userid)
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

func getLectureStatus(lec model.Lecture, now time.Time) string {
	if lec.Finished {
		return "end"
	} else if now.Before(lec.StartAt) {
		return "preparing"
	}
	return "ongoing"
}
