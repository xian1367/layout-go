package timer

import (
	"github.com/go-co-op/gocron"
	"time"
)

var Timer *gocron.Scheduler

func InitTimer() {
	Timer = gocron.NewScheduler(time.UTC)
	Timer.StartAsync()
}

func Shutdown() {
	Timer.Stop()
}
