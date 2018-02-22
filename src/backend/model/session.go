package model

import "time"

//Session 用户登录Token
type Session struct {
	Token       string    `gorm:"type:varchar(255);not null;index;unique_index"`
	UserID      string    `gorm:"type:varchar(20);not null;index"`
	Hint        string    `gorm:"type:varchar(20);not null;default:''"`
	ExpiredTime time.Time `gorm:"type:datetime;not null"`
}
