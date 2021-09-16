package main

import (
	"fmt"
	"time"
)

/**
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
语法
go 函数名( 参数列表 )
*/

func say(s string) {
	for i := 0; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func main() {
	go say("world")
	say("hello")
}
