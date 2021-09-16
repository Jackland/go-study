package main

import "fmt"

/**

切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。

len() 和 cap() 返回结果可相同和不同。


1、直接声明：var slice []int

2、new: slice := *new([]int)

3、字面量：slice := []int{1,2,3,4,5}

4、make: slice :=  make([]int, 5, 10)

5、从切片或数组“截取”:slice := array[1:5] 或 slice := sourceSlice[1:5]
*/

func main()  {

	arr := []int {1,2,3}

	str := []string{"aa","bb"}

	fmt.Println(str)
	fmt.Println(arr)


	a:=make([]int,0)
	//fmt.Println(a[0:1])
	n:=10
	for i:=0;i<=n;i++{
		a = append(a, 1)
		fmt.Printf("len=%d cap=%d\n", len(a), cap(a))
	}
	fmt.Println(a)


	//b:= append(a[0:1], 3)



	fmt.Println(a[0])
	fmt.Println(a[1])

	//fmt.Println(len(b),len(a))
}