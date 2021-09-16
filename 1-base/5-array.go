package main

import "fmt"

func main()  {

	var arr1[5]int
	fmt.Println(arr1)

	arr2:= [6]int{1,2,3,5,6,7}
	fmt.Println(arr2)

	arr3 := [...]int{1,233}
	fmt.Println(arr3)

	//二维数组
	var arr4 = [8][9]int{{1,2},{4,5}}

	arr5:=[2][3]int{{1,3},{3,4}}

	edit(11)

	fmt.Println(arr4)
	fmt.Println(arr5)

}

func edit(a int)  {

	fmt.Println(a)
}