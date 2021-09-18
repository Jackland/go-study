package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counterNum int = 0

func addLock(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counterNum++
	fmt.Printf("%d: %d + %d = %d\n", counterNum, a, b, c)
	lock.Unlock()
}

/***

在 Go 语言中，可供我们选择的同步工具并不少。其中，最重要且最常用的同步工具当属互斥量（mutual exclusion，简称 mutex），sync 包中的 Mutex 就是与其对应的类型，
该类型的值可以被称为互斥锁。一个互斥锁可以被用来保护一个临界区，我们可以通过它来保证在同一时刻只有一个 goroutine 处于该临界区之内，回到我们通过共享内存实现并发通信这篇教程中的示例：

*/
func main() {
	start := time.Now()
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go addLock(1, i, lock)
	}

	for {
		lock.Lock()
		c := counterNum
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consumes := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consumes)
}

//分割线=========================

func addW(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counterNum++
	fmt.Printf("%d: %d + %d = %d\n", counterNum, a, b, c)
	lock.Unlock()
}

func mainW() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go addW(1, i, lock)
	}

	for {
		lock.RLock()
		c := counterNum
		lock.RUnlock()
		runtime.Gosched() //释放缓存
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
