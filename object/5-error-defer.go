package main

import "fmt"

/**
另外，一个函数/方法中可以存在多个 defer 语句，defer 语句的调用顺序遵循先进后出的原则，即最后一个 defer 语句将最先被执行，相当于「栈」结构，如果在循环语句中包含了 defer 语句，
则对应的 defer 语句执行顺序依然符合先进后出的规则。
由于 defer 语句的执行时机和调用顺序，所以我们要尽量在函数/方法的前面定义它们，以免在后面执行时漏掉，尤其是运行时抛出错误会中断后面代码的执行，也就感知不到后面的 defer 语句。
*/
func printError() {
	fmt.Println("兜底执行")
}

func main() {
	defer printError()
	defer func() {
		fmt.Println("除数不能是0!")
	}()

	var i = 1
	var j = 1
	var k = i / j

	fmt.Printf("%d / %d = %d\n", i, j, k)
}
