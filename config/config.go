// Package config 负责配置信息
package config

import (
	"github.com/fsnotify/fsnotify"
	viperLib "github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(configPath string) {
	if _, err := os.Stat(configPath); err != nil {
		panic("配置文件不存在")
	}
	// 初始化 Viper 库
	viper := viperLib.New()
	// 配置类型，支持 "json", "toml", "yaml", "yml", "properties","props", "prop", "env", "dotenv"
	viper.SetConfigType(strings.TrimPrefix(filepath.Ext(configPath), "."))
	// 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(filepath.Dir(configPath))
	// 默认加载 setting.yml 文件，如果有传参 --config=path 的话，加载 path 目录文件
	fileName := filepath.Base(configPath)
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	viper.SetConfigName(fileNameWithoutExt)
	// 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("app")
	// 读取环境变量（支持 flags）
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 配置映射至结构体
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
	// 监控配置文件，变更时重新加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func InitTime() {
	location, err := time.LoadLocation(config.App.Timezone)
	if err != nil {
		panic(err)
	}
	time.Local = location
}

func Get() Config {
	return *config
}
