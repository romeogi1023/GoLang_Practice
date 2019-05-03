package main

import (
	"fmt"
)

//父類
type person04 struct {
	name string
	age  int
	sex  string
}

//父類
type Person04 struct {
	id   int
	addr string
}

//子類
type Student04 struct {
	//結構體成員為多個匿名字段
	Person04 //匿名字段
	person04
	score int
}

func main() {

	var stu Student04
	stu.id = 20
	stu.addr = "北京"
	stu.name = "王富貴"
	stu.age = 10
	stu.sex = "男"
	stu.score = 100
	fmt.Println(stu)

	//自動推導類型
	s := Student04{Person04{200, "北京"}, person04{"王富貴", 10, "男"}, 100}
	fmt.Println(s)
}
