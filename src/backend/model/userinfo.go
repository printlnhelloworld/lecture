package model

import "time"

//UserInfo 用户信息
type UserInfo struct {
	UserID    string    `gorm:"type:varchar(20);primary_key;"`
	Name      string    `gorm:"type:varchar(20);not null;"`
	Type      string    `gorm:"type:varchar(20);not null;"`
	Grade     string    `gorm:"type:int;not null;"`               //年级
	Graduated bool      `gorm:"type:bool;not null;"`              //是否毕业
	Agreed    bool      `gorm:"type:bool;not null;default:false"` //是否同意课外教育规定
	AgreedAt  time.Time `gorm:"type:datetime;"`
	JoinAt    time.Time `gorm:"type:datetime;not null;"`
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
