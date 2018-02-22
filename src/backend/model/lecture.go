package model

import "time"

//Lecture 讲座表
type Lecture struct {
	ID            int       `gorm:"type:int;primary_key;auto_increament;not null"`
	UserID        string    `gorm:"type:varchar(20);index;not null"`
	Topic         string    `gorm:"type:varchar(50);not null;"`
	Introducation string    `gorm:"type:varchar(1000);not null;"`
	StartAt       time.Time `gorm:"type:datetime;not null;"`
	Location      string    `gorm:"type:varchar(100);not null;"`
	Host          string    `gorm:"type:varchar(50);not null;"`
	Lecturer      string    `gorm:"type:varchar(50);not null;"`
	Type          string    `gorm:"type:varchar(50);not null;"`
	Reviewed      string    `gorm:"type:bool;not null;"`
	CreateAt      string    `gorm:"type:datetime;not null;"`
	FinishedAt    string    `gorm:"type:datetime;not null;"`
	Remark        string    `gorm:"type:varchar(100);not null;"`
}

//TODO 讲座类型到底是 int 字典表还是，字符串
