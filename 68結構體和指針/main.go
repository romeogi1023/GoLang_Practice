package main

import (
	"fmt"
)

type person struct {
	id   int
	name string
	age  int
}

func main6801() {
	//初始化
	var per person = person{101, "李寧", 40}
	// fmt.Println(per)
	// fmt.Printf("%p\n", &per)

	//定義指針接收結構體變量地址
	//p := &per
	var p *person = &per
	fmt.Printf("%T\n", p) //*person 類型

	//通過指針間接修改結構體成員的值
	(*p).age = 50
	fmt.Println(per)

	//指針直接修改結構體成員的值
	p.id = 110
	fmt.Println(per)
}

func test68(p *person) {
	p.age = 57
}

func main6802() {
	var per person = person{101, "李寧", 40}
	//地址傳遞
	test68(&per)
	fmt.Println(per)
}

func main03() {
	arr := [3]person{
		{101, "鋼鐵俠", 34},
		{102, "綠巨人", 40},
		{103, "黑寡婦", 28}}
	//指向結構體數組的指針
	p := &arr

	fmt.Printf("%p\n", p)
	fmt.Printf("%T\n", p)

	p[0].age = 40
	fmt.Println(arr[0])
	(*p)[0].age = 39
	fmt.Println(arr[0])

	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}
}

func main() {
	//map類型變量
	m := make(map[int]*[3]person)

	fmt.Printf("%T\n", m)

	m[1] = new([3]person)
	m[1] = &[3]person{{101, "鋼鐵俠", 34},
		{102, "綠巨人", 40},
		{103, "黑寡婦", 28}}

	m[2] = new([3]person)
	m[2] = &[3]person{{101, "美國隊長", 34},
		{102, "黑豹", 40},
		{103, "女巫", 28}}

	for k, v := range m {
		fmt.Println(k, *v)
	}

	//數組z指針
	var p *[3]int

	//創建內存空間 存儲 [3]int
	p = new([3]int)
	p[0] = 123
	p[1] = 456
	p[2] = 789
	fmt.Println(*p)
}
