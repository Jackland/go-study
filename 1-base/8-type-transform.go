package main

import "fmt"

/**

type_name 为类型，expression 为表达式。

*/

func main() {

	var sum int = 12
	var count int = 4
	var mean float32

	mean = float32(sum) / float32(count)

	fmt.Printf("mean的值 ==> %f", mean)

	fmt.Println("--------")

	//不支持隐式转换
	var a int64 = 3
	var b int32
	b = int32(a)
	fmt.Printf("b 为 : %d", b)

}
