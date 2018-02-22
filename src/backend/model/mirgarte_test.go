package model_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
)

//InitDB 初始化数据库
func TestConfTest(t *testing.T) {
	dbconf := conf.Database{
		User:     "root",
		Password: "",
		Addr:     "tcp(192.168.0.101:3306)",
		Db:       "lecture",
	}
	conStr := model.GetDSNFromConf(&dbconf)
	t.Log(conStr)
	_, err := gorm.Open("mysql", conStr)
	if err != nil {
		t.Fatal(err)
	}
}
