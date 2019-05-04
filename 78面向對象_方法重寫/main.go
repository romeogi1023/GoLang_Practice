package main

import (
	"fmt"
)

type person09 struct {
	id   int
	name string
	age  int
}

type student09 struct {
	person09 //匿名字段
	class    int
}

//建議 接收者 使用指針類型 (p *person09)
func (p *person09) PrintInfo() {
	fmt.Printf("編號：%d\n", p.id)
	fmt.Printf("姓名：%s\n", p.name)
	fmt.Printf("年齡：%d\n", p.age)
}

func (s *student09) PrintInfo() {
	fmt.Println("Student：", *s)
}

func main() {
	s := student09{person09{130, "小剛", 18}, 10}
	//子類對象方法 採用就近原則  使用子類的方法
	//方法重寫 子類的方法改寫了父類的方法
	s.PrintInfo()

	//父類對象方法
	s.person09.PrintInfo()

	fmt.Println(s.PrintInfo)
	fmt.Println(s.person09.PrintInfo)
}
