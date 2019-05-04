package main

import (
	"fmt"
)

type cat struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

/*
// 函數定義
// func 函數名(參數列表) 返回值 {
// 	代碼體
// }
func show() {
	fmt.Println("喵喵叫")
}
*/

// 方法定義
// func (對象)方法名(參數列表) 返回值 {
// 	代碼體
// }

func show() {
	fmt.Println("亂亂叫")
}

//結構體類型  可以作為對象類型
//結構體作為接收者
func (c cat) show() {
	fmt.Println("喵喵叫")
	fmt.Printf("我是%s, 喵喵叫\n", c.name)
}

//方法名一樣 接收者不一樣 方法也就不一樣
func (d dog) show() {
	fmt.Println("汪汪叫")
}

func main() {

	//函數調用
	show()

	//對象創建
	var c cat
	c.name = "小花"
	c.age = 2
	fmt.Println(c)
	//方法調用
	c.show()

	var d dog
	d.name = "旺財"
	d.age = 3
	fmt.Println(d)
	d.show()
}
