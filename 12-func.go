package main

import "fmt"

/**

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

type  CalculateType  func(int, int)

func (c *CalculateType ) Server()  {
	fmt.Println("函数类型")
}

func add(a int,b int)  {
	fmt.Println(a+b)
}

func main()  {
	a:=CalculateType(add)
	a(2,4)
	a.Server()
}