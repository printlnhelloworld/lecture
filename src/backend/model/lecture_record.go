package model

//LectureRecord 讲座签到记录表
type LectureRecord struct {
	ID        int    `gorm:"type:int;not null;primary_key;auto_increment"`
	LectureID int    `gorm:"type:int;not null;index"`
	UserID    string `gorm:"type:varchar(20);index;not null"`
	Type      string `gorm:"type:varchar(20);index;not null"`
	CreateAt  string `gorm:"type:datetime;index;not null;"`
	Remark    string `gorm:"type:varchar(20);not null;default ''"`
}
