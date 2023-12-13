package cron

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-module/carbon/v2"
)

func Test() {
	spew.Dump(carbon.Now().ToDateTimeString())
}
