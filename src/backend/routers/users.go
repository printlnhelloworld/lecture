package routers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//GetUsers 获取用户列表 //TODO未实现
func GetUsers() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	}
}

//GetUserByID 获取用户信息
func GetUserByID() func(*gin.Context) {
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
				"stauts": "ok",
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

//GetLecturesByUserID 获取用户的参加的讲座列表
func GetLecturesByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//GetLectureByLectureIDByUserID 获取用户的参加的特定讲座
func GetLectureByLectureIDByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//AddTokensByUserID 登录
func AddTokensByUserID() func(*gin.Context) {
	return func(c *gin.Context) {
		userid := c.Param("userid")
		type login struct {
			ID       string `json:"id" binding:"required"`
			Password string `json:"password" binding:"required"`
			Remark   string `json:"remark"`
		}
		var loginvar login
		c.ShouldBindWith(&loginvar, binding.JSON)
		if loginvar.ID == "" || loginvar.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "id/password必须不为空",
			})
			return
		}
		if userid != loginvar.ID {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ParamErr",
				"msg":    "userid 和 id 不相同",
			})
			return
		}

	}
}

//UserLoginCallBack //登录回调
func UserLoginCallBack(appconf *conf.Conf) func(*gin.Context) {
	return func(c *gin.Context) {
		ticket := c.Query("ticket")
		service := appconf.BaseURL + "/api/v1/loginCallback"
		encodeURL := "http://cas.hdu.edu.cn/cas/serviceValidate?ticket=" + ticket + "&service=" + url.QueryEscape(service)
		resp, err := http.Get(encodeURL)
		if err != nil {
			c.Redirect(http.StatusFound, "?auth=&err=CasGetErr"+err.Error())
		} else {
			databytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c.Redirect(http.StatusFound, "?auth=&err=ReadErr"+err.Error())
			} else {
				m := ParseUserInfoFromCas(string(databytes))
				if len(m) == 0 {
					c.JSON(http.StatusUnauthorized, gin.H{
						"status": "NoDataErr",
						"msg":    "没有数据，请重新登陆",
					})
					return
				}
				c.JSON(http.StatusOK, m)
			}
		}
	}
}

//GetLoginURL 获取登录连接，给前端使用
func GetLoginURL(appconf *conf.Conf) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"msg":      "ok",
			"loginURL": "http://cas.hdu.edu.cn/cas/login?service=" + url.QueryEscape(appconf.BaseURL+"/api/v1/loginCallback"),
		})
	}
}

//GetTokensByUserID 登录列表
func GetTokensByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

//DeleteTokenByUserID 登出
func DeleteTokenByUserID() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

var re = regexp.MustCompile(`<sso:attribute name="(.*)" type="java.lang.String" value="(.*)"/>`)

//ParseUserInfoFromCas 解析数据
func ParseUserInfoFromCas(data string) (m map[string]string) {
	m = map[string]string{}
	for _, match := range re.FindAllStringSubmatch(data, -1) {
		m[match[1]] = match[2]
	}
	return
}
