package main

import (
	"fmt"
)

/**
用法详解
1,Go语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，也就是说，
先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行

2, 关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源，典型的例子就是对一个互斥解锁，或者关闭一个文件
*/

func main() {
	fmt.Println("defer begin")

	defer fmt.Println(1)

	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println("defer end")

}

/**
defer 修饰的 f.Close() 方法会在函数执行完成后或读取文件过程中抛出错误时执行，以确保已经打开的文件资源被关闭，从而避免内存泄露。如果一条语句干不完清理的工作，
也可以在 defer 后加一个匿名函数来执行对应的兜底逻辑：
*/
