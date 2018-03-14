package model

import "time"

//Permit 权限
type Permit struct {
	Permit   string    `json:"permit" gorm:"type:varchar(100);not null;primary_key;"`
	UserID   string    `json:"userid" gorm:"type:varchar(20);not null;primary_key;"`
	CreateAt time.Time `json:"createAt" gorm:"datetime;not null"`
}

//GetUserPermits 获取用户权限
func GetUserPermits(uid string) []string {
	var permits []Permit
	DB.Find(&permits, "user_id = ?", uid)
	var r []string
	for _, p := range permits {
		r = append(r, p.Permit)
	}
	return r
}

//UpdateUserPermits 更新用户权限
func UpdateUserPermits(uid string, ps []string) error {
	//TODO
	return nil
}
