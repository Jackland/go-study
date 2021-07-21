package main

import "fmt"

//全局变量
var hh,jj int




var (
	ll int
	pp string
	cc bool

)

func main()  {




	var a string = "Jayson"
	var c,d int = 1,6
	fmt.Println("a==",a)
	fmt.Println("c=",c,"d=",d)

	h:=111;
	fmt.Println("h=",h)
	fmt.Println(hh,jj)

	fmt.Println("ll=,pp=,cc=",ll,pp,cc)

	//函数返回多个值
	mm,numb,strs := nums()
	fmt.Println(mm,numb,strs)


	_,numbs,strss := nums() //只获取函数返回值的后两个
	fmt.Println(numbs,strss)

}


//一个函数返回多个值
func nums()(int,int,string){
	a, b, c := 1, 2, "gg"
	return a, b, c;
}