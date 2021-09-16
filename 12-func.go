package main

import (
	"errors"
	"fmt"
)

/**
和 php 的区别
Go 语言的函数和 PHP 函数的基本结构差不多，只是具体细节有所差异，
在 Go 语言中，函数的基本组成为：关键字 func、函数名、参数列表、返回值、函数体和返回语句，作为强类型语言，无论是参数还是返回值，在定义函数时，都要声明其类型。



go语言函数的表达式
func 函数名（形参名1 类型1，形参名2 类型2，…）（返回类型1，返回类型2，…） {
函数体
}
注意：go语言的函数支持多个返回值。函数名大写表示该函数为Public可以被其它package调用，小写为private，不可以被其它包调用。

ex1:
func myFunc(number1 int32,number2 int32) (int32,int32) {
  return number1+number2 ,number1 -number2
}

*/

/**
ex01 下面我们以 add01 函数为例介绍函数的定义，如果函数的参数列表中包含若干个类型相同的参数，比如上面例子中的 a 和 b，则可以在参数列表中省略前面变量的类型声明
同时要声明返回类型 int
*/

func add01(a, b int) int {
	return a + b
}

/**
ex02 函数的直接调用
在跨包调用时，只有首字母大写的函数才可以被调用，这个涉及到包的可见性，后面在介绍面向对象时会详细讨论这个，这里你只需要了解首字母大写的 Go
函数即相当于其它语言的 public 函数，Go 语言中没有 public、protected、private 之类的关键字，它是通过首字母的大小写来区分可见性：
首字母小写的函数只能在同一个包中访问，首字母大写的函数才可以在其他包中调用，Go 文件中定义的全局变量也是如此。
*/

/**
go 值传递和地址传递
*/

//1，值传递
func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		err := errors.New("只能大于0相加")
		return 0, err
	}

	a *= 2 //乘法
	b *= 3
	return a + b, nil
}

//2.地址传递
func addArea(a, b *int) (int, int, int) {
	*a *= 2 //乘法
	*b *= 3
	return *a, *b, *a + *b
}

//3.返回其他类型的数据
func fn() (res []string) {
	for i := 0; i < 10; i++ {
		res = append(res, "A")
	}
	return
}

func fn1() (res []string) {
	res = make([]string, 10)
	for i := 0; i < 10; i++ {
		res[i] = "A"
	}
	return
}

/**
go的一下内置函数 不需要导入任何包 close,len、cap,new、make,copy、append...

前面我们提到 new 函数作用于值类型，仅分配内存空间，返回的是指针，make 函数作用于引用类型，除了分配内存空间，还会对对应类型进行初始化，
返回的是初始值。在 Go 语言中，引用类型包括切片（slice）、字典（map）和管道（channel），其它都是值类型。


*/

func main() {
	//4.匿名函数
	acc := func(a, b int) int {
		return a + b
	}
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)
	fmt.Println(acc(1, 2)) // 调用匿名函数 acc

	return

	fmt.Print(fn())
	fmt.Print(fn1())

	return
	x, y := -1, 2
	z, err := add(x, y)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//_,_,z := addArea(&x, &y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
}

type CalculateType func(int, int)

func (c *CalculateType) Server() {
	fmt.Println("函数类型")
}

//func add(a int,b int)  {
//	fmt.Println(a+b)
//}

//func main()  {
//	a:=CalculateType(add)
//	a(2,4)
//	a.Server()
//}
