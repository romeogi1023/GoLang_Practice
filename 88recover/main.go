package main

import (
	"fmt"
)

func testA() {
	fmt.Println("testA")
}

func testB(x int) {

	//設置recover()
	//在defer調用的函數中使用recover()
	defer func() {

		//防止程序崩潰
		//recover()

		//防止程序崩潰 並印出錯誤訊息
		//fmt.Println(recover())

		//防止程序崩潰 並且有錯誤時才打印錯誤訊息
		if err := recover(); err != nil {
			fmt.Println(err)
		}

	}() //匿名函數

	var a [3]int
	a[x] = 999
	fmt.Println(a[x])
}

func testC() {
	fmt.Println("testC")
}

func main() {
	testA()
	testB(3) //發生異常	中斷程序
	testC()
}
