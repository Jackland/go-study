package main

import "fmt"

//全局变量
var hh, jj int

//定义一个常量
const Pi float64 = 3.14159265358979323846

const ( // 通过一个 const 关键字定义多个常量，和 var 类似
	size int64 = 1024
	eof        = -1 // 无类型整型常量
)

var (
	ll int
	pp string
	cc bool
)

func main() {
	fmt.Println(Pi, size)
	//缺省参数的写法
	_, nickname := GetName()
	fmt.Println(nickname)

	fmt.Println("===================")
	var a string = "Jayson"
	var c, d int = 1, 6
	fmt.Println("a==", a)
	fmt.Println("c=", c, "d=", d)

	h := 111
	fmt.Println("h=", h)
	fmt.Println(hh, jj)

	fmt.Println("ll=,pp=,cc=", ll, pp, cc)

	//函数返回多个值
	mm, numb, strs := nums()
	fmt.Println(mm, numb, strs)

	_, numbs, strss := nums() //只获取函数返回值的后两个
	fmt.Println(numbs, strss)

}

func GetName() (userName, nickName string) {
	return "nonfu", "学院君"
}

//一个函数返回多个值
func nums() (int, int, string) {
	a, b, c := 1, 2, "gg"
	return a, b, c
}
