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
		userid := c.Param("userid")
		u, err := model.GetUserByID(userid)
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
					"permit": map[string]string{
						"": "",
					},
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

//GetUserLectureByLectureID 获取用户的参加的特定讲座
func GetUserLectureByLectureID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//GetUserTokens 登录列表
func GetUserTokens() func(*gin.Context) {
	return func(c *gin.Context) {
		userid, exist := c.Get("UserID")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "UnauthErr",
				"msg":    "未认证",
			})
			return
		}
		tokens := model.GetUserTokensByUserID(userid.(string))
		m := []map[string]interface{}{}
		for _, t := range *tokens {
			m = append(m, map[string]interface{}{
				"remark":   t.Remark,
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
		token, exist := c.Get("Token")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "UnauthErr",
				"msg":    "未认证",
			})
			return
		}

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
		userid, exist1 := c.Get("UserID")
		token, exist2 := c.Get("Token")
		if !exist1 || !exist2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "UnauthErr",
				"msg":    "未认证",
			})
			return
		}
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
