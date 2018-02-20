package model

import "time"

//Announcement 公告
type Announcement struct {
	ID         int       `gorm:"type:int;primary_key;auto_increment" json:"id"`
	UserID     string    `gorm:"type:varchar(20);" json:"userId"`
	UserName   string    `gorm"-" json:"userName"`
	Title      string    `gorm:"type:varchar(50);"`
	Content    string    `gorm:"type:varchar(1000);"`
	Important  bool      `gorm:"type:bool"`
	CreateTime time.Time `gorm:"type:datetime"`
}

func getAnnouncements() {

}
