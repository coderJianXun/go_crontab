package worker

import (
	"encoding/json"
	"io/ioutil"
)

var (
	G_config *Config
)

// 程序配置
type Config struct {
}

func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	// 把配置文件读进来
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// JSOn反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 赋值单例
	G_config = &conf

	return
}
