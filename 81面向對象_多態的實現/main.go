package main

import (
	"fmt"
)

//先定義接口  一般以er結尾  根據接口實現功能
type Humaner1 interface {
	//方法
	sayhi()
}

type student81 struct {
	name  string
	age   int
	score int
}

func (s *student81) sayhi() {
	fmt.Printf("大家好，我是%s，今年%d歳，我的成績%d分\n", s.name, s.age, s.score)
}

type teacher81 struct {
	name    string
	age     int
	subject string
}

func (t *teacher81) sayhi() {
	fmt.Printf("大家好，我是%s，今年%d歳，我的學科%s\n", t.name, t.age, t.subject)
}

//多態的實現
//將接口作為函數參數  實現多態
func sayhello(h Humaner1) {
	h.sayhi()
}

func main() {

	stu := student81{"小明", 18, 98}
	//調用多態函數
	sayhello(&stu)

	tea := teacher81{"老王", 28, "物理"}
	sayhello(&tea)
}
