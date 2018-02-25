package conf_test

import (
	"testing"

	"git.hduhelp.com/hduhelper/lecture/src/backend/conf"
)

func TestConf(t *testing.T) {
	c, err := conf.LoadConfig("./app.toml.example")
	if err != nil {
		t.Error(err)
	}
	t.Log(c.App)
	t.Log(c.ListenAddr)
	t.Log(c.BaseURL)
	t.Log(c.Database)
}
