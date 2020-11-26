package main

import (
	"github.com/geiqin/gotools/helper"
	"github.com/geiqin/gotools/xcode"
	"github.com/geiqin/gotools/xtime"
	"github.com/shomali11/util/xhashes"
	log "log"
)



func main() {
	log.Println("code:", helper.GenerateSn())
	log.Println("code:", helper.GenerateSn("2018"))
	c :=xcode.GetRand()
	log.Println(c)
	xhashes.FNV64("ddd")
	xtime.ParseTimeToDate("2020-11-12 32:12")
	log.Println("dddd")
}
