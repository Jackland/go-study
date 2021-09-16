package main

import (
	"errors"
	"fmt"
)

/**
语法
type error interface {
	Error() string
}

*/

func Foo(param int) (n int, err error) {
	// ...
	if param > 2 {
		err = errors.New("error happen")
		return
	}
	return n, err
}

func main() {
	n, err := Foo(2323)
	if err != nil {
		fmt.Println(err)
		// 错误处理
	} else {
		// 使用返回值 n
		fmt.Println(n)
	}

}
