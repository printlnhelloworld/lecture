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
func InitDB(conf *conf.Conf) error {
	dbconf := conf.Database
	conStr := dbconf.User + ":" + dbconf.Password + "@" + dbconf.Addr + "/" + dbconf.Db
	conStr += "?charset=utf8mb4"
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		return errors.New(conStr + err.Error())
	}

	db.AutoMigrate(Announcement{})
	DB = db
	return nil
}
