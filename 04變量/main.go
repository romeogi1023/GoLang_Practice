package main

import (
	"fmt"
)

func main() {
	//一、變量的聲明
	//1. 聲明 var 變量名 類型    變量聲明後，必須使用
	//2. 只是聲明變 默認為0
	//3. 在同一個{}裡，聲明的變量是唯一的

	var a int
	fmt.Println(a)

	// var b, c int
	// b, c = 20, 30
	// fmt.Println(b, c)

	//二、變量的初始化  聲明變量時，同時賦值
	var b int = 10 //初始化  一步到位
	b = 20         //賦值  先聲明 再賦值
	fmt.Println(b)

	var c float64 = 3.14
	fmt.Println(c)
	var value float64 = 2
	//跟表達式
	var sum float64 = value * value
	fmt.Println(sum)
}
