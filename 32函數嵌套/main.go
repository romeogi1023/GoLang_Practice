package main

import "fmt"

func test1(num1, num2 int) {
	fmt.Println(num1 + num2)
}

func test(a, b int) {
	test1(a, b)
}

func test3(arr ...int) {
	fmt.Println(arr)
}

func test2(arr ...int) {

	//不能將不定參的名稱傳遞給另1個不定參
	//test3(arr...int)  //error

	//傳遞指定個數數據
	//test3(arr[0],arr[1],arr[2],arr[3])
	//傳遞多個數據
	test3(arr[0:5]...)

	//傳遞全部參數
	//test3(arr[:]...)
	//test3(arr...)

}

func main() {

	//1.什麽是函數的嵌套
	//在1個函數中調用另1個函數
	//2.函數嵌套的執行過程
	//test(10,20)
	//3.函數嵌套調用的應用

	//4.不定參數函數調用
	test2(1, 2, 3, 4, 5, 6)

}
