package main

/**
go 语言中 启动多携程 难点在于并发引起的不确定性，以及为了协调这种不确定性在不同协程间所要进行的通信。

常见的并发通信模型有两种：共享内存和消息传递。
*/

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func add(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

/**
这种写法 代码也因此变得更复杂，更难以维护，这还只是个简单的加法运算实现，就要写这么多代码，要引入共享变量，还要引入互斥锁来保证操作的原子性，
对于更加复杂的业务代码，如果到处都要加锁、解锁，显然对开发者和维护者来说都是噩梦，Go 语言既然以并发编程作为语言的核心优势，当然不至于将这样的问题用这么繁琐的方式来解决。

解决方案看 3-channel.go
*/
func main() {
	start := time.Now()
	lock := &sync.Mutex{} //锁机制，即 sync.Mutex，这是 Go 语言标准库提供的互斥锁
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
