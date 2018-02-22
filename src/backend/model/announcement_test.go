package model_test

import (
	"testing"
	"time"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
)

func init() {
	model.InitDB(&conf.Database{
		User:     "root",
		Password: "",
		Addr:     "tcp(192.168.0.101:3306)",
		Db:       "lecture",
	})
}

var t = time.Date(2018, 2, 13, 15, 16, 20, 0, time.FixedZone("Beijing Time", int((8*time.Hour).Seconds())))
var data = []model.Announcement{
	{
		ID:        1,
		UserID:    "15051237",
		Title:     "维护通知",
		Content:   "xx年x月x日 x时x分开始维护，请做好准备",
		Important: true,
		CreateAt:  t,
	},
	{
		ID:        2,
		UserID:    "15051237",
		Title:     "维护通知",
		Content:   "2018年2月30日 1时1分开始维护，请做好准备",
		Important: true,
		CreateAt:  t,
	},
	{
		ID:        3,
		UserID:    "15051236",
		Title:     "新版本上线通知",
		Content:   "2018年2月31日 上线新版本，本次新增加xxx",
		Important: false,
		CreateAt:  t,
	},
}

func TestCreateAnnoucement(t *testing.T) {
	for _, d := range data {
		aid, err := model.CreateAnnouncement(d.Title, d.Content, d.UserID, d.Important)
		if err != nil {
			t.Error(aid, err)
		}
	}
}

func TestGetAnnoucementByID(t *testing.T) {
	for i, d := range data {
		a, err := model.GetAnnouncementByID(i + 1)
		if err != nil {
			t.Error(err)
		} else if a.ID != d.ID {
			t.Error("ID error")
		} else if a.Content != d.Content {
			t.Error("ContentError")
		} else if a.UserID != d.UserID {
			t.Error("UserIDError")
		} else if a.Title != d.Title {
			t.Error("Title Error")
		}
	}
}

func TestDeleteAnnoucementByID(t *testing.T) {
	for i, _ := range data {
		err := model.DeleteAnnouncementByID(i + 1)
		if err != nil {
			t.Error(err)
		}
	}
}
