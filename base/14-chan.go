package main

import (
	"fmt"
	"time"
)

/**
通道声明和初始化
通过 chan 类型关键字来声明通道类型变量
var chs [10]chan int
var chs []chan int
var chs map[string]chan int

//实际编码 还可以传递第二个参数，表示通道的容量：
ch:=make(chan int,10)

//通道修饰符
通道类型变量只支持发送和接收操作，即往通道中写入数据和从通道中读取数据，对应的操作符都是 <-，
我们判断是发送还是接收操作的依据是通道类型变量位于 <- 左侧还是右侧，位于左侧是发送操作，位于右侧是接收操作：

ch <- 1   // 往通道中写入数据 1
x := <- ch     // 从通道中读取数据并赋值给变量


//2-我们还可以通过如下方式实现双向通道和单向通道的转化
1
ch1 := make(chan int)
2
ch2 := <-chan int(ch1)//读（接收通道）
3
ch3 := chan<- int(ch1) //写(发送通道)

*/

func test(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	//调用 close(ch) 显式关闭通道，这一行不能漏掉，否则主协程不知道子协程什么时候执行完毕，从一个空的通道接收数据会报如下运行时错误（死锁）：
	close(ch)
}

func main() {
	start := time.Now()
	ch := make(chan int, 10)
	go test(ch)

	for i := range ch {
		fmt.Println("数据接收：", i)
	}
	end := time.Now()
	//耗时
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
