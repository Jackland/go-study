package main

import "fmt"

/**
for init; condition; post { }


 */

func main()  {
	sum:= 0
	for i:=0;i<=10;i++ {
		sum+=i
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum<=10 {
		sum+=sum
	}
	fmt.Println(sum)

	//a:=0
	//无限循环 无法输出信息内容
	//for {
	//	a++
	//}
	//fmt.Println(a)


}