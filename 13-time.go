package main

import (
	"fmt"
	"time"
)

/**
go 里面时间相关函数
*/

func main() {
	//1, 表示1s  time.Second * 1 （golang的sleep不是系统调用，只是协程层面的挂起。）
	time.Sleep(time.Second * 1) //睡眠1秒
	time.Sleep(1)               // 会睡眠 1 纳秒，时间很短

	//测试
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)

}
