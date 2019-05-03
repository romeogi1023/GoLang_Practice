package main

import (
	"fmt"
)

//結構體嵌套結構體
//父類
type Person struct {
	id   int
	name string
	age  int
}

//子類
type Student struct {
	Person //匿名字段
	score  int
}

func main7201() {

	//對象創建
	//順序初始化
	var s1 Student = Student{Person{101, "小明", 18}, 98}
	fmt.Println(s1)

	// var p1 Person = Person{102, "小亮", 28}
	// fmt.Println(p1)

	//自動推導類型
	s2 := Student{Person{101, "小亮", 28}, 87}
	fmt.Println(s2)

	//指定初始化成員 沒有初始化的部份 使用默認值
	s3 := Student{score: 97}
	fmt.Println(s3)

	s4 := Student{Person{name: "小紅"}, 87}
	fmt.Println(s4)

}

func main() {
	var s1 Student = Student{Person{101, "小明", 18}, 98}
	s1.score = 89
	s1.Person.age = 20
	s1.age = 30

	s1.Person = Person{110, "小蘭", 20}
	fmt.Println(s1)
}
