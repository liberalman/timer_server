package main

import (
	"log"
	"runtime"
	"timer_server/timer"
)

func callback1(args interface{}) {
	//只执行一次的事件
	if values, ok := args.([]string); ok {
		var str1 string = values[0]
		var str2 string = values[1]
		log.Println("callback1(" + str1 + "," + str2 + ")")
	} else {
		log.Println("callback1()")
	}
}

func callback2(args interface{}) {
	//每次在当前时间点之后5s插入一个定时器，这样就能形成每隔5秒调用一次callback2回调函数，可以用于周期性事件
	timer.SetTimer("callback2", 5, callback2, args)
	log.Println("callback2")
}

func main() {
	// cpu多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 定时器1，传入两个参数
	timer.SetTimer("callback1", 3, callback1, []string{"hello", "world"})
	// 定时器2，不传参数
	timer.SetTimer("callback2", 6, callback2, nil)
	// 移除定时器
	//timer.Delete(timer.TimerMap["callback2"])
	//运行计时器
	timer.Run()
}
