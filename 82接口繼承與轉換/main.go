package main

import (
	"fmt"
)

//先定義接口  一般以er結尾  根據接口實現功能
type Humaner2 interface { //子集
	//方法
	sayhi()
}

type Personer interface { //超集
	Humaner2 //匿名字段
	sing(string)
}

type student82 struct {
	name  string
	age   int
	score int
}

func (s *student82) sayhi() {
	fmt.Printf("大家好，我是%s，今年%d歳，我的成績%d分\n", s.name, s.age, s.score)
}

func (s *student82) sing(name string) {
	fmt.Println("我為大家唱首歌", name)
}

type teacher82 struct {
	name    string
	age     int
	subject string
}

func (t *teacher82) sayhi() {
	fmt.Printf("大家好，我是%s，今年%d歳，我的學科%s\n", t.name, t.age, t.subject)
}

//多態的實現
//將接口作為函數參數  實現多態
func sayhello(h Humaner2) {
	h.sayhi()
}

func main82() {

	//接口類型變量定義
	var h Humaner2
	stu := student82{"小明", 18, 98}
	h = &stu
	h.sayhi()

	//接口類型變量定義
	var p Personer
	p = &stu
	p.sayhi()
	p.sing("大碗面")
}

func main() {
	//接口類型變量定義
	var h Humaner2 //子集
	var p Personer //超集
	stu := student82{"小明", 18, 98}
	p = &stu
	//將一個接口賦值給另一個接口
	//超集中包含所子集的方法
	h = p

	//子集不包含超集
	//不能將子集賦值給超集
	//p = h	//err

	h.sayhi()
}
