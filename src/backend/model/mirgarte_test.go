package model_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
)

//InitDB 初始化数据库
func TestConfTest(t *testing.T) {
	dbconf := conf.Database{
		"root",
		"",
		"tcp(192.168.0.101:3306)",
		"lecture",
	}
	conStr := dbconf.User + ":" + dbconf.Password + "@" + dbconf.Addr + "/" + dbconf.Db
	conStr += "?charset=utf8mb4"
	t.Log(conStr)
	_, err := gorm.Open("mysql", conStr)
	if err != nil {
		t.Fatal(err)
	}
}
