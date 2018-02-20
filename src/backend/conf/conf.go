package conf

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

//Conf 配置文件
type Conf struct {
	App        string   `toml:"app"`        //App 名称
	ListenAddr string   `toml:"listenAddr"` //监听地址如 0.0.0.0:8000
	Database   database `toml:"database"`   //数据库配置
}

type database struct {
	User     string `toml:"user"`     //用户名
	Password string `toml:"password"` //密码
	Addr     string `toml:"addr"`     //数据库地址 tcp(ip:port)
	Db       string `toml:"db"`       //默认数据库
}

//LoadConfig 加载配置
func LoadConfig(path string) (*Conf, error) {
	var conf Conf
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
