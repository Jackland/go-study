package main

import (
	"fmt"
	"sync"
)

/**

sync.WaitGroup 类型是开箱即用的，也是并发安全的。该类型提供了以下三个方法：

Add：WaitGroup 类型有一个计数器，默认值是0，我们可以通过 Add 方法来增加这个计数器的值，通常我们可以通过个方法来标记需要等待的子协程数量；
Done：当某个子协程执行完毕后，可以通过 Done 方法标记已完成，该方法会将所属 WaitGroup 类型实例计数器值减一，通常可以通过 defer 语句来调用它；
Wait：Wait 方法的作用是阻塞当前协程，直到对应 WaitGroup 类型实例的计数器值归零，如果在该方法被调用的时候，对应计数器的值已经是 0，那么它将不会做任何事情。

*/

func addNum(a, b int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go addNum(i, 1, wg.Done)
	}
	wg.Wait()
}
