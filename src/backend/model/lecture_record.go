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

//GetLectureRecordsByLectureID 根据讲座 id 获取签到记录
func GetLectureRecordsByLectureID(lid int) {

}

//GetLectureRecordsByUserID 根据用户 id 获取签到记录
func GetLectureRecordsByUserID(uid string) {

}

//AddLectureRecord 添加讲座签到记录
func AddLectureRecord(uid string, lid int) {

}

//DeleteLectureRecord 删除签到记录
func DeleteLectureRecord(lid int, uid string) {

}

//DeleteLectureAllRecord 删除特定讲座的所有签到记录
func DeleteLectureAllRecord(lid int) {

}
