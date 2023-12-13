package bootstrap

import (
	"github.com/xian137/layout-go/config"
	"github.com/xian137/layout-go/pkg/logger"
	"github.com/xian137/layout-go/pkg/timer"
)

// SetupLogger 初始化 Logger
func SetupLogger(label string) {
	log(label)
	_, _ = timer.Timer.Every(1).Day().At("00:00:00").Do(log)
}

func log(label string) {
	logger.InitLogger(
		config.Get().Log.MaxSize,
		config.Get().Log.MaxBackup,
		config.Get().Log.MaxAge,
		config.Get().Log.Level,
		config.Get().Log.FilePath+label+"-",
	)
}
