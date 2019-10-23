package main

import (
	_ "log"
	"runtime"
	"time"

	"github.com/xuexihuang/goHeartTimer/timer"
)

var socketcount int = 2

type TimeArg struct {
	socket int
	time   uint32
	count  int
}

func callback1(args interface{}) {
	timeargs := args.(TimeArg)
	//log.Printf("this is socket:%d", timeargs.socket)
	timeargs.count++
	if timeargs.count > 1000 {

	} else {

		timer.SetTimer("netheart", timeargs.time, callback1, timeargs)
	}
}
func SetHeart() {
	for index := 0; index < 20000; index++ {
		timeargs := TimeArg{index, 20, 1}
		timer.SetTimer("netheart", 20, callback1, timeargs)
		time.Sleep(1 * time.Millisecond)
	}
}
func main() {
	// cpu多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	go SetHeart()
	timer.Run()
}
