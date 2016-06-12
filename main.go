// timingwheel project main.go
package main

import (
	"log"
	"runtime"
	"timer_server/timer"
)

//全局变量定义
var (
	//timerMap      map[string]*util.Node = make(map[string]*util.Node) //保存待执行的计时器，方便按链表节点指针地址直接删除定时器
	funcName map[string]func(interface{})
)

type TimerData struct {
	exectime int64       //到期执行的时间,(unix时间戳格式)
	funcName string      //函数名称
	args     interface{} //函数参数
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func callback1(args interface{}) {
	log.Println("callback1")
}

func callback2(args interface{}) {
	//每隔60秒
	timer.SetTimer("callback2", 5, callback2, args)
	log.Println("callback2")
}

func main() {
	// cpu多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	timer.SetTimer("callback1", 1, callback1, nil)
	timer.SetTimer("callback2", 5, callback2, nil)
	//运行计时器，间隔1s
	timer.Run()
}
