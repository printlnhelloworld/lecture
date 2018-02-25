package model_test

import (
	"fmt"
	"testing"

	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
)

func TestParseCas(t *testing.T) {
	var str = []string{`<?xml version='1.0' encoding='UTF-8'?>
<sso:serviceResponse xmlns:sso="sso-namespace">
  <sso:authenticationSuccess>
    <sso:user>15051237</sso:user>
    <sso:attributes>
      <sso:attribute name="user_name" type="java.lang.String" value="杨飞"/>
      <sso:attribute name="id_type" type="java.lang.String" value="1"/>
      <sso:attribute name="userName" type="java.lang.String" value="15051237"/>
      <sso:attribute name="user_id" type="java.lang.String" value="201508173964"/>
      <sso:attribute name="unit_id" type="java.lang.String" value="27"/>
      <sso:attribute name="user_sex" type="java.lang.String" value="1"/>
      <sso:attribute name="unit_name" type="java.lang.String" value="网络空间安全学院、浙江保密学院"/>
      <sso:attribute name="classid" type="java.lang.String" value="15052411"/>
    </sso:attributes>
  </sso:authenticationSuccess>
</sso:serviceResponse>`, `
<?xml version='1.0' encoding='UTF-8'?>
<sso:serviceResponse xmlns:sso="sso-namespace">
  <sso:authenticationSuccess>
    <sso:user>07033</sso:user>
    <sso:attributes>
      <sso:attribute name="user_name" type="java.lang.String" value="胡海滨"/>
      <sso:attribute name="id_type" type="java.lang.String" value="3"/>
      <sso:attribute name="userName" type="java.lang.String" value="07033"/>
      <sso:attribute name="user_id" type="java.lang.String" value="201306005378"/>
      <sso:attribute name="unit_id" type="java.lang.String" value="20"/>
      <sso:attribute name="user_sex" type="java.lang.String" value="1"/>
      <sso:attribute name="unit_name" type="java.lang.String" value="材料与环境工程学院"/>
      <sso:attribute name="classid" type="java.lang.String" value=""/>
    </sso:attributes>
  </sso:authenticationSuccess>
</sso:serviceResponse>
`,
	}
	for _, s := range str {
		m := model.ParseUserInfoFromCas(s)
		fmt.Println(m)
	}
}
