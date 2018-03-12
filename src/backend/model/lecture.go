package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Lecture 讲座表
type Lecture struct {
	ID               int       `gorm:"type:int AUTO_INCREMENT;primary_key;not null"` //讲座id
	UserID           string    `gorm:"type:varchar(20);index;not null"`              //创建者
	Topic            string    `gorm:"type:varchar(50);not null;"`                   //主题
	Introduction     string    `gorm:"type:text;not null;"`                          //简介
	StartAt          time.Time `gorm:"type:datetime;index;not null;"`                //开始时间
	Location         string    `gorm:"type:varchar(100);not null;"`                  //地点
	Host             string    `gorm:"type:varchar(50);not null;"`                   //主办方
	Lecturer         string    `gorm:"type:varchar(50);not null;"`                   //主讲人
	Type             string    `gorm:"type:varchar(50);index;not null;"`             //讲座类型
	Reviewed         bool      `gorm:"type:bool;index;not null;"`                    //是否同意讲座
	ReviewedBy       string    `gorm:"type:varchar(20);not null;"`                   //同意讲座的人
	CanSignin        bool      `gorm:"type:bool;not null;"`                          //讲座是否开始签到
	SignCode         string    `gorm:"type:varchar(20);not null;default:''"`         //签到码
	SignCodeExpireAt time.Time `gorm:"type:datetime;not null;"`                      //签到码过期时间
	Finished         bool      `gorm:"type:bool;index;not null;"`                    //讲座是否结束
	FinishedAt       time.Time `gorm:"type:datetime;not null;"`                      //讲座结束时间

	CreateAt time.Time `gorm:"type:datetime;not null;"`
	Remark   string    `gorm:"type:varchar(100);not null;"`
}

//TODO 讲座类型到底是 int 字典表还是，字符串

//GetLectures 获取讲座列表
func GetLectures(limit, next int, owner, status, sort string, now time.Time) (*[]Lecture, error) {
	//TODO 添加 reviewed 参数，来表示是否通过审核
	var db *gorm.DB
	db = DB.Order("finished asc")
	switch sort {
	case "id":
		db = db.Order("id desc")
		if next != 0 {
			db = db.Where("id < ?", next)
		}
	case "startAt":
		db = db.Order("start_at desc")
		if next != 0 {
			db = db.Where("start_at < ?", next)
		}
	}
	if owner != "" {
		db = db.Where("user_id = ?", owner)
	}
	nowF := now.Format("2006-01-02 15:04:05")
	switch status {
	case "prepare":
		db = db.Where("start_at > ?", nowF)
	case "runing":
		db = db.Where("start_at <= ? and finished_at > ?", nowF, nowF)
	case "ended":
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
func CreateLecture(userid, topic, location, introduction, host, lecturer, lectype string, reviewed bool, startAt time.Time) (int, error) {
	var lec Lecture
	lec.UserID = userid
	lec.Topic = topic
	lec.Introduction = introduction
	lec.StartAt = startAt
	lec.Location = location
	lec.Host = host
	lec.Lecturer = lecturer
	lec.Type = lectype
	lec.Reviewed = reviewed
	lec.CreateAt = time.Now()
	return lec.ID, DB.Create(&lec).Error
}

//PatchLecture 修改讲座
func PatchLecture(lid int, m map[string]interface{}) error {
	return DB.Model(&Lecture{}).Where(&Lecture{ID: lid}).Update(m).Error
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
	_ = DeleteLectureAllRecord(id)
	return db.Error
}

//UpdateLectureStatus 更新讲座状态
func UpdateLectureStatus(lid int, status string) error {
	var err error
	switch status {
	case "signing":
		err = DB.Model(Lecture{ID: lid}).Update("can_signin", "1").Error
	case "notsigning":
		err = DB.Model(Lecture{ID: lid}).Update("can_signin", "0").Error
	case "ended":
		err = DB.Model(Lecture{ID: lid}).Update(Lecture{Finished: true, FinishedAt: time.Now()}).Error
	}
	return err
}

//UpdateLectureSignCode 获取讲座签到码
func UpdateLectureSignCode(lid int, code string, t time.Time) error {
	lec := Lecture{
		ID: lid,
	}
	return DB.Model(&lec).Update(Lecture{SignCode: code, SignCodeExpireAt: t}).Error
}
