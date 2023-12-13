package cron

import (
	"github.com/xian137/layout-go/pkg/console"
	"github.com/xian137/layout-go/pkg/timer"
)

func Kernel() {
	var err error
	_, err = timer.Timer.Every(1).Second().Do(Test)
	console.ExitIf(err)
}
