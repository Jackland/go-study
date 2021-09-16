package main

import "fmt"

/**
Map 集合是无序的 key-value 数据结构。

Map 集合中的 key / value 可以是任意类型，但所有的 key 必须属于同一数据类型，所有的 value 必须属于同一数据类型，key 和 value 的数据类型可以不相同
 */
func main()  {

	var p  map[int]string
	p = make(map[int]string)

	p[0] = "my"
	p[1] = "mine"
	p[2] = "me"

	var map1 map[string]int = map[string]int{}
	map1["aaa"] = 111
	fmt.Println(map1)

	var map2 = map[int]string{}
	map2[1] = "ab"
	map2[2] = "ccc"
	fmt.Println(map2)

	map3:= map[int]string{
		1:"tom",
		2:"jack",
	}
	fmt.Println(map3)

}