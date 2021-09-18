package main

import "fmt"

/**

1、init函数用于包的初始化，如初始化包中的变量，这个初始化在package xxx的时候完成，也就是在main之前完成；

2、每个包可以拥有多个init函数， 每个包的源文件也可以拥有多个init函数；

3、同一个包中多个init函数的执行顺序是没有明确定义的，但是不同包的init函数是根据包导入的依赖关系决定的。

4、init函数不能被其他函数调用，其实在main函数之前自动执行的。

*/

func main() {
	fmt.Println("do in main")
}

func init() {
	fmt.Println("do in init1")
}

func init() {
	fmt.Println("do in init2")
}
