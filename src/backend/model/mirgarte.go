package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
)

//DB 数据库连接
var DB *gorm.DB

//InitDB 初始化数据库
func InitDB(dbconf *conf.Database) error {
	conStr := GetDSNFromConf(dbconf)
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		return errors.New(conStr + err.Error())
	}

	db.LogMode(true)
	db.AutoMigrate(
		&Announcement{},
		&Lecture{},
		&LectureRecord{},
		&UserInfo{},
		&Token{},
		&Permit{},
	)
	DB = db
	return nil
}

//GetDSNFromConf 从配置中拼接DSN字符串
func GetDSNFromConf(dbconf *conf.Database) string {
	conStr := dbconf.User + ":" + dbconf.Password + "@" + dbconf.Addr + "/" + dbconf.Db
	conStr += "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	return conStr
}

//TODO 更新数据库设计表
