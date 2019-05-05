package main

import "fmt"

func main() {
	//1.聲明變量 沒有初始化  零值 為false
	var a bool
	fmt.Println(a)

	a = false
	fmt.Println(a)
	//2.布爾類型不接受其他類型的賦值，不支持自動或強制的類型轉換
	//a = 1
	//a = bool(1)
	//fmt.Println(a)

	//3.自動推導類型

	var b = true
	fmt.Println(b)

	c := false
	fmt.Println(c)

	v2 := (1 == 2)
	fmt.Println(v2)
	fmt.Printf("%T", v2)

}
