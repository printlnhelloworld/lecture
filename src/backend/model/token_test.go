package model_test

import "testing"

import (
	"git.hduhelp.com/hduhelper/lecture/src/backend/model"
)

func TestGetUserIDByToken(t *testing.T) {
	id, err := model.GetUserIDByToken("token")
	if err == nil {
		t.Error(id)
	}
	if err = model.AddToken("111111", "token", "127.0.0.1"); err != nil {
		t.Error(err.Error())
	}

	id, err = model.GetUserIDByToken("token")
	if err != nil {
		t.Error(err.Error())
	}
	if id != "111111" {
		t.Error(id)
	}
	if err = model.DeleteToken("self", "111111", "token"); err != nil {
		t.Error(err.Error())
	}
}
