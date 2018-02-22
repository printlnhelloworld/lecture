package routers

import (
	"net/http"
	"strconv"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

type annnouncement struct {
	ID         int    `json:"id"`
	Important  bool   `json:"important" binding:"required"`
	Author     string `json:"author"`
	AuthorName string `json:"authorName"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CreateAt   int64  `json:"createAt"`
}

//GetAnnouncements 获取所有公告
func GetAnnouncements() func(*gin.Context) {
	return func(c *gin.Context) {
		next, err1 := strconv.Atoi(c.DefaultQuery("next", "0"))
		limit, err2 := strconv.Atoi(c.DefaultQuery("limit", "10"))

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "param type error",
				"msg":    "参数类型错误 limit/next 必须为int",
			})
			return
		}
		if limit > 20 {
			limit = 20
		}
		anns := model.GetAnnouncements(next, limit)

		tmp := make([]annnouncement, 0)
		for _, ann := range *anns {
			tmp = append(tmp, *conv(&ann))
		}

		newNext := 0
		if len(tmp) > 0 {
			newNext = tmp[len(tmp)-1].ID
		}
		resp := gin.H{
			"status": "ok",
			"msg":    "ok",
			"next":   newNext,
			"count":  len(tmp),
			"data":   tmp,
		}
		c.JSON(http.StatusOK, resp)
	}
}

//CreateAnnouncements 创建公告
func CreateAnnouncements() func(*gin.Context) {
	return func(c *gin.Context) {
		var a annnouncement
		if err := c.ShouldBindWith(&a, binding.JSON); err == nil {
			userid, exist := c.Get("UserID")
			if !exist {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "ServerError",
					"msg":    "服务出现错误",
					"err":    "Lost UserID",
				})
			}
			if aid, err := model.CreateAnnouncement(a.Title, a.Content, userid.(string), a.Important); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "ServerError",
					"msg":    "服务出现错误",
					"err":    "CreateError",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
					"msg":    "ok",
					"id":     aid,
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamLostError",
				"msg":    "必须包含title、important、content",
				"err":    err.Error(),
			})
		}
	}
}

//GetAnnouncementByID 获取单个公告
func GetAnnouncementByID() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("announcementid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "param type error",
				"msg":    "参数类型错误",
				"next":   c.Param("announcementid"),
			})
			return
		}
		ann, err := model.GetAnnouncementByID(id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "ErrRecordNotFound",
					"msg":    "没有数据",
				})
			} else {
				c.JSON(http.StatusBadGateway, gin.H{
					"status": "databaseError",
					"msg":    "数据库错误",
					"err":    err.Error(),
				})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"data":   *conv(ann),
		})
	}
}

//PutAnnouncementByID 修改公告
func PutAnnouncementByID() func(*gin.Context) {
	return func(c *gin.Context) {
		aid, err := strconv.Atoi(c.Param("announcementid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamError",
				"msg":    "参数错误",
				"err":    "announcementid must be int",
			})
			return
		}
		var a annnouncement
		if err := c.ShouldBindWith(&a, binding.JSON); err == nil {
			userid, exist := c.Get("UserID")
			if !exist {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "ServerError",
					"msg":    "服务出现错误",
					"err":    "Lost UserID",
				})
			} else if err := model.PutAnnouncement(aid, a.Title, a.Content, userid.(string), a.Important); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "ServerError",
					"msg":    "服务出现错误",
					"err":    err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
					"msg":    "ok",
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamLostError",
				"msg":    "必须包含title、important、content",
				"err":    err.Error(),
			})
		}
	}
}

//DeleteAnnouncementByID 删除公告
func DeleteAnnouncementByID() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("announcementid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "param type error",
				"msg":    "参数类型错误",
				"next":   c.Param("announcementid"),
			})
			return
		}
		count, err := model.DeleteAnnouncementByID(id)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"status": "databaseError",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
		} else if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "RecordNotFound",
				"msg":    "记录不存在",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"msg":    "ok",
			})
		}
	}
}

func conv(ann *model.Announcement) *annnouncement {
	a := annnouncement{}
	a.ID = ann.ID
	a.Author = ann.UserID
	a.AuthorName = ann.Author.Name
	a.Title = ann.Title
	a.Important = ann.Important
	a.Content = ann.Content
	a.CreateAt = ann.CreateAt.Unix()
	return &a
}
