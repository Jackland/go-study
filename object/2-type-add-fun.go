package main

import "fmt"

/**
在 Go 语言中，你可以给任意类型（包括基本类型，但不包括指针类型）添加成员方法，但是如果是基本类型的话，需要借助 type 关键字对类型进行再定义;
*/

type Integer int

/**
相当于在类里面添加方法 封装对应功能点
*/
func (a Integer) Equal(b Integer) bool {
	return a == b
}

func (a Integer) LessThan(b Integer) bool {
	return a < b
}

func (a Integer) MoreThan(b Integer) bool {
	return a > b
}

func (a *Integer) increase(i Integer) {
	*a = *a + i
}

func (a *Integer) decrease(i Integer) {
	*a = *a - i
}

func main() {
	var a Integer = 2
	if a.Equal(2) {
		fmt.Println(a, "is equal to 2")
	}

	a.increase(12)
	fmt.Println(a)
	a.decrease(12)
	fmt.Println(a)

}
