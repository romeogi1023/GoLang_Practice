package main

import "fmt"

//定義函數
func run() {
	fmt.Println("跑")
}
func walk() {
	fmt.Println("走")
}

func main() {

	//一.什麽是函數？
	//函數就是將壹些代碼進行重用的壹種機制

	//二.函數的基本語法

	//func 函數名(參數列表) 返回值 {
	//   函數體
	// }

	//1.func關鍵字
	//2.函數名不能重名   圓括號  花括號

	//三.函數的使用
	run()
	walk()
	//函數名+（）

}
