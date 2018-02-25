package model

import "time"

//Token 用户登录Token
type Token struct {
	Token    string    `gorm:"type:varchar(100);not null;primary_key"`
	UserID   string    `gorm:"type:varchar(20);not null;index"`
	Hint     string    `gorm:"type:varchar(20);not null;default:''"`
	CreateAt time.Time `gorm:"type:datetime;not null"` //创建时间
	ExpireAt time.Time `gorm:"type:datetime;not null"` //过期时间
}

//AddToken 添加用户token
func AddToken(userid, token string) error {
	return DB.Create(
		&Token{
			Token:    token,
			UserID:   userid,
			Hint:     "",
			CreateAt: time.Now(),
			ExpireAt: time.Now().AddDate(0, 0, 30),
		},
	).Error
}

//DeleteToken 删除用户token
func DeleteToken(token string) error {
	return DB.Delete(&Token{Token: token}).Error
}

//GetUserIDByToken 根据token获取用户id
func GetUserIDByToken(token string) (string, error) {
	t := Token{
		Token: token,
	}
	err := DB.Find(&t).Error
	if err != nil {
		return "", err
	} else {
		return t.UserID, nil
	}
}
