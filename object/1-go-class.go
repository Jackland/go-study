package main

import "fmt"

/**

类的定义和初始化
Go 语言的面向对象编程与我们之前所熟悉的 PHP、Java 那一套完全不同，没有 class、extends、implements 之类的关键字和相应的概念，而是借助结构体来实现类的声明，
比如要定义一个学生类，可以这么做：
*/
type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

/**
类名为 Student，并且包含了 id、name、male、score 四个属性，Go 语言中也不支持构造函数、析构函数，
取而代之地，可以通过定义形如 NewXXX 这样的全局函数（首字母大写）作为类的初始化函数：
*/
func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

//为类添加成员方法
func (s Student) getName() string {
	return s.name
}

//修改类的属性  Student 类型设置成了指针类型
func (s *Student) setName(name string) {
	s.name = name
}

func main() {
	student := NewStudent(1, "jayson", false, 100)

	student.setName("land")

	fmt.Println(student.name)
	//调用成员方法
	fmt.Println("name:", student.getName())
	fmt.Println(student)
}
