package main

import "fmt"

var Usage = func() {
	fmt.Println("USAGE: Hello usage [arguments] ...")
	fmt.Println("\n这是第一行\n 这是第二行")
}

func main() {
	Usage()
	fmt.Println("hello,go")
}
