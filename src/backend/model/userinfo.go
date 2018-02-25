package model

import (
	"time"
)

//UserInfo 用户信息
type UserInfo struct {
	UserID   string    `gorm:"type:varchar(20);primary_key;"`         //学号/工号
	Name     string    `gorm:"type:varchar(20);not null;"`            //姓名
	Type     string    `gorm:"type:varchar(20);not null;"`            //类别
	ClassID  string    `gorm:"type:varchar(20);not null; default:''"` //班级号
	Sex      bool      `gorm:"type:bool;not null;"`                   //性别
	UnitID   string    `gorm:"type:varchar(20);not null;"`            //学院编号
	UnitName string    `gorm:"type:varchar(100);not null;"`           //学院名称
	Agreed   bool      `gorm:"type:bool;not null;default:false"`      //是否同意课外教育规定
	AgreedAt time.Time `gorm:"type:datetime;"`                        //同意课外规定时间
	JoinAt   time.Time `gorm:"type:datetime;not null;"`               //第一次进入本系统时间
}

//TODO 是否需要isatshcool isregister，这个在 api.hdu 中有。在cas中没有

//UserCas Cas中数据的结构
type UserCas struct {
	UserName string
	IDType   string
	UserID   string
	UnitID   string
	UserSex  string
	UnitName string
	ClassID  string
}

//GetUserByID 通过ID查用户信息
func GetUserByID(id string) (*UserInfo, error) {
	var user UserInfo
	DB.Find(&user, id)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return &user, nil
}
