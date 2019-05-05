package main

import "fmt"

func main() {
	//byte字符類型  同時也uint8的別名
	//1.單引號 字符 雙引號  字符串
	var a byte = 'a'
	fmt.Println(a)
	//%c是一個占位符  表示打印輸出一個字符
	fmt.Printf("%c\n", 97)
	fmt.Printf("%T\n", a)

	var b byte = '0'
	fmt.Println(b)

	//\n \t  \0 字符串的結束的標誌

	fmt.Println('\n')

}
