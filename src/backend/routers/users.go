package routers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
	"git.hduhelp.com/hduhelper/lecture/src/backend/model"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
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

//UserLoginCallBack //登录回调
func UserLoginCallBack(appconf *conf.Conf) func(*gin.Context) {
	return func(c *gin.Context) {
		ticket := c.Query("ticket")
		service := appconf.BaseURL + "/api/v1/loginCallback"
		encodeURL := "http://cas.hdu.edu.cn/cas/serviceValidate?ticket=" + ticket + "&service=" + url.QueryEscape(service)
		baseHashURL := "/app/#/login?auth="
		resp, err := http.Get(encodeURL)
		if err != nil {
			c.Header("Location", baseHashURL+"&err=CasGetErr&msg="+err.Error()+"&msg1=服务错误")
		} else {
			databytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c.Header("Location", baseHashURL+"&err=ReadErr"+"&msg="+err.Error())
			} else {
				m := ParseUserInfoFromCas(string(databytes))
				if len(m) == 0 {
					c.Header("Location", baseHashURL+"&err=UnauthErr&msg=登录出现错误，请重试")
				} else {
					if err := model.UpdateUserInfo(m); err != nil {
						c.Header("Location", baseHashURL+"&err=DatabaseErr&msg="+err.Error()+"&msg1=数据库错误")
					}
					token := rendToken()
					if err := model.AddToken(m["userName"], token); err != nil {
						c.Header("Location", baseHashURL+"&err=DatabaseErr&msg="+err.Error()+"&msg1=数据库错误")
					} else {
						c.Header("Location", baseHashURL+token)
					}
				}
			}
		}
		c.JSON(http.StatusFound, gin.H{})
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
func GetUserTokens() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
	}
}

//DeleteUserToken 登出
func DeleteUserToken() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "ok",
			"msg":    "ok",
		})
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

func rendToken() string {
	u := uuid.NewV4()
	return u.String()
}
