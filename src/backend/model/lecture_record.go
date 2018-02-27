package model

import "time"

//LectureRecord 讲座签到记录表
type LectureRecord struct {
	LectureID int       `gorm:"type:int;primary_key"`
	UserID    string    `gorm:"type:varchar(20);primary_key"`
	Type      string    `gorm:"type:varchar(20);not null"`
	CreateAt  time.Time `gorm:"type:datetime;index;not null;"`
	Remark    string    `gorm:"type:varchar(20);not null;default ''"`
}

//GetLectureRecordsByLectureID 根据讲座 id 获取签到记录
func GetLectureRecordsByLectureID(lid int) (total int, lrs []LectureRecord) {
	DB.Order("`create_at` desc").Where("`lecture_id` = ?", lid).Find(&lrs)
	return len(lrs), lrs
}

//GetLectureRecordsByUserID 根据用户 id 获取签到记录
func GetLectureRecordsByUserID(uid string) []LectureRecord {
	var lrs []LectureRecord
	DB.Find(&lrs, "`user_id` = ?", uid)
	return lrs
}

//AddLectureRecord 添加讲座签到记录
func AddLectureRecord(rtype, uid string, lid int) (LectureRecord, error) {
	lr := LectureRecord{
		LectureID: lid,
		UserID:    uid,
		Type:      rtype,
		CreateAt:  time.Now(),
		Remark:    "",
	}
	err := DB.Create(&lr).Error
	return lr, err
}

//DeleteLectureRecord 删除签到记录
func DeleteLectureRecord(lid int, uid string) error {
	return DB.Delete(LectureRecord{}, "`lecture_id` = ? AND `user_id` = ? AND `type` = 'byhand'", lid, uid).Error
}

//DeleteLectureAllRecord 删除特定讲座的所有签到记录
func DeleteLectureAllRecord(lid int) error {
	return DB.Delete(LectureRecord{}, "`lecture_id` = ?", lid).Error
}
