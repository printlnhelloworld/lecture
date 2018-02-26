package model

import "time"

//Token 用户登录Token
type Token struct {
	Token    string    `gorm:"type:varchar(100);not null;primary_key"`
	UserID   string    `gorm:"type:varchar(20);not null;index"`
	IP       string    `gorm:"type:varchar(25);not null;"` //用户登录时ip
	Remark   string    `gorm:"type:varchar(20);not null;default:''"`
	CreateAt time.Time `gorm:"type:datetime;not null"` //创建时间
	ExpireAt time.Time `gorm:"type:datetime;not null"` //过期时间
}

//AddToken 添加用户token
func AddToken(userid, token, ip string) error {
	return DB.Create(
		&Token{
			Token:    token,
			UserID:   userid,
			Remark:   "",
			IP:       ip,
			CreateAt: time.Now(),
			ExpireAt: time.Now().AddDate(0, 0, 30),
		},
	).Error
}

//DeleteToken 删除用户token
func DeleteToken(filter, userid, token string) error {
	switch filter {
	case "all":
		return DB.Delete(Token{}, "`user_id` = ?", userid).Error
	case "self":
		return DB.Delete(&Token{Token: token}).Error
	case "other":
		return DB.Delete(Token{}, "`user_id` = ? AND `token` <> ?", userid, token).Error
	}
	return nil
}

//GetUserIDByToken 根据token获取用户id
func GetUserIDByToken(token string) (string, error) {
	t := Token{
		Token: token,
	}
	err := DB.Where("`expire_at` >= ?", time.Now().Format("2006-01-02 15:04:05")).Find(&t).Error
	if err != nil {
		return "", err
	}
	return t.UserID, nil
}

//GetUserTokensByUserID 通过用户 id 获取 token列表
func GetUserTokensByUserID(userid string) *[]Token {
	tokens := make([]Token, 0)
	DB.Where(&Token{UserID: userid}).Where("`expire_at` >= ?", time.Now().Format("2006-01-02 15:04:05")).Find(&tokens)
	return &tokens
}

//UpdateTokenRemark 更新 Token 备注
func UpdateTokenRemark(token, remark string) error {
	return DB.Model(&Token{}).Where(Token{Token: token}).Update("remark", remark).Error
}
