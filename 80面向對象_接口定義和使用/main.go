package main

import (
	"fmt"
)

//先定義接口  一般以er結尾  根據接口實現功能
type Humaner interface {
	//方法
	sayhi()
}

type student80 struct {
	name  string
	age   int
	score int
}

func (s *student80) sayhi() {
	fmt.Printf("大家好，我是%s，今年%d歳，我的成績%d分\n", s.name, s.age, s.score)
}

type teacher80 struct {
	name    string
	age     int
	subject string
}

func (t *teacher80) sayhi() {
	fmt.Printf("大家好，我是%s，今年%d歳，我的學科%s\n", t.name, t.age, t.subject)
}

func main() {

	//接口是一種數據類型 可以接收滿足對象的信息
	//接口是虛的  方法是實的
	//接口定義規則  方法實現規則
	//接口定義的規則 在方法中必須有定義的實現
	var h Humaner

	stu := student80{"小明", 18, 98}
	//stu.sayhi()
	//將對象信息賦值給接口類型變量
	h = &stu
	h.sayhi()

	tea := teacher80{"老王", 28, "物理"}
	//tea.sayhi()
	//將對象賦值給接口 滿足接口中的方法的聲明格式
	h = &tea
	h.sayhi()
}
