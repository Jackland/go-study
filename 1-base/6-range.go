package main

import "fmt"

//Go 语言中
//range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

/**
for key, value := range oldMap {
    newMap[key] = value
}

 */


func main()  {
	nums:=[]int{1,2,3}
	sum:=0

	// _ 是key 可省略不写
	for _,num:=range nums{
		sum+=num
	}

	for a,num:=range  nums{
		sum+=num
		if num == 3{
			fmt.Println("index",num)
		}
		fmt.Println(a)
	}
	fmt.Println(sum)

	fmt.Println("--------------------")
	//range 也可以用于 map 上
	maps :=map[string]string{"a":"jack","b":"rose"}

	for a,v:=range maps{
		fmt.Printf("key%s-val%s\n",a,v)
	}

	fmt.Println("-----------------")
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}