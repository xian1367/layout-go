package cron

import (
	"github.com/xian1367/layout-go/pkg/console"
	"github.com/xian1367/layout-go/pkg/timer"
)

func Kernel() {
	var err error
	_, err = timer.Timer.Every(1).Second().Do(Test)
	console.ExitIf(err)
}
