package routers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"

	"github.com/gin-gonic/gin"
)

//GetLectures 获取讲座列表
func GetLectures() func(*gin.Context) {
	//TODO 接口中添加排序方式 id / startAt
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
	return func(c *gin.Context) {

	}
}

//PutLectureByID 修改讲座，不用带上全部参数
func PutLectureByID() func(*gin.Context) {
	return func(c *gin.Context) {

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
						"creatorName":   "", //TODO 未实现
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

	}
}

//GetSigninRecordLecturesByID 获取特定讲座签到记录
func GetSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//DeleteOneSigninRecordLecturesByID 删除特定讲座签到记录
func DeleteOneSigninRecordLecturesByID() func(*gin.Context) {
	return func(c *gin.Context) {

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
