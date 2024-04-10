// Package app 应用信息
package app

import (
	"github.com/xian1367/layout-go/config"
)

func IsLocal() bool {
	return config.Get().App.Mode == "local"
}

func IsTesting() bool {
	return config.Get().App.Mode == "test"
}

func IsStage() bool {
	return config.Get().App.Mode == "stage"
}

func IsProduction() bool {
	return config.Get().App.Mode == "production"
}

func IsDebug() bool {
	return config.Get().App.Debug
}
