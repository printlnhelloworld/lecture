package model

import (
	"time"
)

//UserInfo 用户信息
type UserInfo struct {
	UserID   string     `gorm:"type:varchar(20);primary_key;"`         //学号/工号
	Name     string     `gorm:"type:varchar(20);not null;"`            //姓名
	Type     string     `gorm:"type:varchar(20);not null;"`            //类别
	ClassID  string     `gorm:"type:varchar(20);not null; default:''"` //班级号
	Sex      string     `gorm:"type:varchar(1);not null;"`             //性别
	UnitID   string     `gorm:"type:varchar(20);not null;"`            //学院编号
	UnitName string     `gorm:"type:varchar(100);not null;"`           //学院名称
	Agreed   bool       `gorm:"type:bool;not null;default:false"`      //是否同意课外教育规定
	AgreedAt *time.Time `gorm:"type:datetime;not null;"`               //同意课外规定时间
	JoinAt   *time.Time `gorm:"type:datetime;not null;"`               //第一次进入本系统时间
}

//TODO 是否需要isatshcool isregister，这个在 api.hdu 中有。在cas中没有
//TODO 移除 unitname 到其他表中

//GetUserByID 通过ID查用户信息
func GetUserByID(id string) (*UserInfo, error) {
	var user UserInfo
	DB.Find(&user, id)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return &user, nil
}

//UpdateUserInfo 更新用户信息
func UpdateUserInfo(m map[string]string) error {
	//TODO 可能有数据竞争
	var user UserInfo
	now := time.Now()
	err := DB.Find(&user, m["userName"]).Error
	if err != nil {
		user = UserInfo{
			UserID:   m["userName"],
			Name:     m["user_name"],
			Type:     m["id_type"],
			Sex:      m["user_sex"],
			UnitID:   m["unit_id"],
			UnitName: m["unit_name"],
			ClassID:  m["classid"],
			Agreed:   false,
			AgreedAt: &now,
			JoinAt:   &now,
		}
		err = DB.Create(&user).Error
	} else {
		user = UserInfo{
			UserID:   m["userName"],
			Name:     m["user_name"],
			Type:     m["id_type"],
			Sex:      m["user_sex"],
			UnitID:   m["unit_id"],
			UnitName: m["unit_name"],
			ClassID:  m["classid"],
		}
		err = DB.Update(&user).Error
	}
	return err
}
