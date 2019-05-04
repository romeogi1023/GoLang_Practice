package main

import (
	"fmt"
)

type person08 struct {
	id   int
	name string
	age  int
}

type student08 struct {
	person08 //匿名字段
	class    int
}

//建議 接收者 使用指針類型 (p *person08)
func (p *person08) PrintInfo() {
	fmt.Printf("編號：%d\n", p.id)
	fmt.Printf("姓名：%s\n", p.name)
	fmt.Printf("年齡：%d\n", p.age)
}

func main() {
	p := person08{110, "德瑪", 30}
	p.PrintInfo()

	//子類可以繼承父類 同時也繼承屬性和方法
	//父類不能繼承子類 不論是屬性和方法
	s := student08{person08{120, "德邦", 15}, 9}
	s.PrintInfo()
}
