package model

import (
	"time"
)

//Announcement 公告
type Announcement struct {
	ID        int       `gorm:"type:int AUTO_INCREMENT;primary_key" json:"id"` //TODO 好像不太好设置自动增加
	UserID    string    `gorm:"type:varchar(20);NOT NULL;"`
	Author    UserInfo  `gorm:"foreignkey:UserID;association_foreignkey:UserID" json:"author"`
	Title     string    `gorm:"type:varchar(50);NOT NULL" json:"title"`
	Content   string    `gorm:"type:varchar(1000);NOT NULL" json:"content"`
	Important bool      `gorm:"type:bool;index;NOT NULL" json:"important"`
	CreateAt  time.Time `gorm:"type:datetime;index;NOT NULL"`
}

//GetAnnouncements 获取所有公告
func GetAnnouncements(next, limit int) *[]Announcement {
	anns := make([]Announcement, 0)
	if next == 0 {
		DB.Order("id desc").Limit(limit).Find(&anns)
	} else {
		DB.Order("id desc").Where("id < ?", next).Limit(limit).Find(&anns)
	}
	for i := range anns {
		//TODO 好像性能差不多 都是 5~10ms
		DB.Find(&anns[i].Author, anns[i].UserID)
		//DB.Model(&anns[i]).Related(&anns[i].Author, "UserID")
	}
	return &anns
}

//CreateAnnouncement 保存公告
func CreateAnnouncement(title, content, userid string, important bool) (id int, err error) {
	var a = Announcement{
		Title:     title,
		Content:   content,
		UserID:    userid,
		Important: important,
		CreateAt:  time.Now(),
	}

	if err := DB.Create(&a).Error; err != nil {
		return 0, err
	}
	return a.ID, nil
}

//PutAnnouncement 修改公告
func PutAnnouncement(aid int, title, content, userid string, important bool) error {
	var ann Announcement
	if err := DB.Find(&ann, aid).Error; err != nil {
		return err
	}
	ann.Title = title
	ann.Content = content
	ann.Important = important
	ann.UserID = userid

	return DB.Save(&ann).Error
}

//DeleteAnnouncementByID 删除公告
func DeleteAnnouncementByID(id int) (int64, error) {
	db := DB.Delete(Announcement{}, "id = ?", id)
	return db.RowsAffected, db.Error
}

//GetAnnouncementByID 通过ID获取公告
func GetAnnouncementByID(id int) (*Announcement, error) {
	var a Announcement
	if err := DB.Find(&a, id).Error; err != nil {
		return nil, err
	}

	DB.Model(&a).Related(&a.Author, "UserID")
	return &a, nil

}
