package main

import (
	"fmt"
)

//爺
type humen struct {
	id   int
	name string
}

//父
type person05 struct {
	humen //匿名字段
	age   int
	sex   string
}

//子
type student05 struct {
	person05 //匿名字段
	score    int
}

func main() {
	var stu student05
	stu.name = "亞索"
	stu.sex = "男"
	stu.score = -5
	fmt.Println(stu)

	//初始化
	//自動推導類型
	stu1 := student05{person05{humen{1001, "亞索"}, 30, "男"}, -5}
	fmt.Println(stu1)
}
