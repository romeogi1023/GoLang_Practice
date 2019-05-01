package main

import (
	"fmt"
)

//結構體定義在函數外部
// type 結構體名 struct {
// 	//結構體成員列表
// 	//成員名 數據類型
// 	name string
// 	age  int
// }

//結構體名   大小寫需注意   若要給其它檔案調用 則第一個字要大寫
type student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main5801() {

	//順序初始化 每個成員必須初始化
	var s1 student = student{1, "小明", "男", 18, "市區"}
	fmt.Println(s1)

	//自動推導類型
	stu := student{name: "小花", age: 18, sex: "女", id: 2, addr: "市區"}
	fmt.Println(stu)

	//定義結構體變量  複合類型
	//var 變量名 結構體名
	var s2 student
	//賦值
	//為結構體成員賦值  . 包名.函數名  結構體.成員  對象.方法
	s2.name = "小亮"
	s2.age = 19
	s2.sex = "男"
	s2.id = 3
	s2.addr = "郊區"
	fmt.Println(s2)
}

func main5802() {
	stu := student{1, "小紅", "女", 19, "市區"}
	//結構體名本身指向第一個成員的地址
	fmt.Printf("%p\n", &stu)
	fmt.Printf("%p\n", &stu.id)
	fmt.Printf("%p\n", &stu.name)
}

func main() {
	// s1 := student{1, "小明", "男", 18, "市區"}
	// s2 := student{1, "小明", "男", 18, "市區"}
	// s3 := student{2, "小亮", "男", 18, "市區"}
	// fmt.Println(s1 == s2)
	// fmt.Println(s1 == s3)

	//結構體賦值
	s3 := student{2, "小亮", "男", 18, "市區"}

	var s4 student
	s4 = s3
	fmt.Println(s4)
}
