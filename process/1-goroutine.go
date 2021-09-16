package main

import (
	"fmt"
	"time"
)

func addA(a, b int) {
	var c = a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	for i := 0; i < 10; i++ {
		go addA(1, i)
	}
	time.Sleep(3e9) //睡眠1秒
}
