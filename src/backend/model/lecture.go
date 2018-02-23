package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Lecture 讲座表
type Lecture struct {
	ID            int       `gorm:"type:int AUTO_INCREMENT;primary_key;not null"`
	UserID        string    `gorm:"type:varchar(20);index;not null"`
	Topic         string    `gorm:"type:varchar(50);not null;"`
	Introducation string    `gorm:"type:varchar(1000);not null;"`
	StartAt       time.Time `gorm:"type:datetime;index;not null;"`
	Location      string    `gorm:"type:varchar(100);not null;"`
	Host          string    `gorm:"type:varchar(50);not null;"`
	Lecturer      string    `gorm:"type:varchar(50);not null;"`
	Type          string    `gorm:"type:varchar(50);index;not null;"`
	Reviewed      bool      `gorm:"type:bool;index;not null;"`
	Finished      bool      `gorm:"type:bool;index;not null;"`
	CreateAt      time.Time `gorm:"type:datetime;not null;"`
	FinishedAt    time.Time `gorm:"type:datetime;not null;"`
	Remark        string    `gorm:"type:varchar(100);not null;"`
}

//TODO 讲座类型到底是 int 字典表还是，字符串

//GetLectures 获取讲座列表
func GetLectures(limit, next int, owner, status, sort string, now time.Time) (*[]Lecture, error) {
	var db *gorm.DB
	switch sort {
	case "id":
		db = DB.Order("id desc")
		if next != 0 {
			db = db.Where("id < ?", next)
		}
	case "startAt":
		db = DB.Order("start_at desc")
		if next != 0 {
			db = db.Where("start_at < ?", next)
		}
	}
	if owner != "" {
		db = db.Where("user_id = ?", owner)
	}
	nowF := now.Format("2006-01-02 15:04:05")
	switch status {
	case "preparing":
		db = db.Where("start_at > ?", nowF)
	case "ongoing":
		db = db.Where("start_at <= ? and finished_at > ?", nowF, nowF)
	case "end":
		db = db.Where("finished = ?", true)
	default:
	}
	if limit != 0 {
		db = db.Limit(limit)
	}
	var lectures []Lecture
	err := db.Find(&lectures).Error
	if err != nil {
		return nil, err
	}
	return &lectures, nil
}

//CreateLecture 创建讲座
func CreateLecture() error {
	return nil
}

//PatchLecture 修改讲座
func PatchLecture(lid int) error {
	return nil
}

//GetLectureByID 获取特定讲座信息
func GetLectureByID(lid int) (*Lecture, error) {
	var lec Lecture
	err := DB.Find(&lec, lid).Error
	return &lec, err
}

//DeleteLectureByID 删除讲座
func DeleteLectureByID(id int) error {
	db := DB.Delete(&Lecture{}, id)
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Error
}
