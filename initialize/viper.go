package initialize

import (
	"fmt"
	"github.com/akazwz/go-resume/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitViper 初始化配置
func InitViper() (config *viper.Viper) {
	config = viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	// 读取配置
	if err := config.ReadInConfig(); err != nil {
		panic(err)
		return nil
	}

	// 解析配置
	if err := config.Unmarshal(&global.CFG); err != nil {
		panic(err)
		return nil
	}

	// 监听配置
	config.WatchConfig()
	// 热修改配置
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file updated:", e.Name)
		if err := config.Unmarshal(&global.CFG); err != nil {
			panic(err)
		}
	})
	return
}
