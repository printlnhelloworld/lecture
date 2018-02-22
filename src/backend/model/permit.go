package model

import "time"

//Permit 权限
type Permit struct {
	ID       int       `json:"id" gorm:"primary_key;not null;"`
	Scope    int       `json:"scope" gorm:"not null;index"`
	UserID   string    `json:"userid" gorm:"not null';index"`
	CreateAt time.Time `json:"createAt" gorm:"not null"`
}
