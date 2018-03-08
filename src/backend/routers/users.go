package routers

import (
	"net/http"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//GetUserInfo 获取用户信息
func GetUserInfo() func(*gin.Context) {
	return func(c *gin.Context) {
		userid, _ := c.Get("UserID")

		u, err := model.GetUserByID(userid.(string))
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"status": "DatabaseError",
				"msg":    "数据库错误",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"msg":    "ok",
				"data": map[string]interface{}{
					"id":       u.UserID,
					"name":     u.Name,
					"type":     u.Type,
					"classId":  u.ClassID,
					"sex":      u.Sex,
					"unitID":   u.UnitID,
					"unitName": u.UnitName,
					"agree":    u.Agreed,
					"agreeAt":  u.AgreedAt.Unix(),
					"joinAt":   u.JoinAt.Unix(),
					"permit":   model.GetUserPermits(u.UserID),
				},
			})
		}
	}
}

//GetUserLectures 获取用户的参加的讲座列表
func GetUserLectures() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//UpdateUserAgree 更新用户同意用户协议（包含课外教育部分）
func UpdateUserAgree() func(*gin.Context) {
	return func(c *gin.Context) {
		userid, _ := c.Get("UserID")
		err := model.UpdateUserAgree(userid.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "datbaseErr",
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

//GetUserTokens 登录列表
func GetUserTokens() func(*gin.Context) {
	return func(c *gin.Context) {
		userid, _ := c.Get("UserID")

		tokens := model.GetUserTokensByUserID(userid.(string))
		m := []gin.H{}
		for _, t := range *tokens {
			m = append(m, gin.H{
				"remark":   t.Remark,
				"ip":       t.IP,
				"createAt": t.CreateAt.Unix(),
				"expireAt": t.ExpireAt.Unix(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
			"list":   m,
		})
	}
}

//UpdateUserTokenRemark 更新 token 备注
func UpdateUserTokenRemark() func(*gin.Context) {
	return func(c *gin.Context) {
		token, _ := c.Get("Token")

		type remark struct {
			Remark string `json:"remark"`
		}
		rmk := remark{}
		if err := c.ShouldBindWith(&rmk, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "参数错误",
				"err":    err.Error(),
			})
			return
		}
		err := model.UpdateTokenRemark(token.(string), rmk.Remark)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "databaseErr",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	}
}

//DeleteUserToken 登出
func DeleteUserToken(filter string) func(*gin.Context) {
	return func(c *gin.Context) {
		userid, _ := c.Get("UserID")
		token, _ := c.Get("Token")

		err := model.DeleteToken(filter, userid.(string), token.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "databaseErr",
				"msg":    "数据库错误",
				"err":    err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	}
}
