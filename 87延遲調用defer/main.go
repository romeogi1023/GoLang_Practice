package main

import "fmt"

func test87(x int) {
	v := 100 / x
	fmt.Println(v)
}

func main() {
	/*
		//defer語句只能出現在函數內部
		defer fmt.Println("hello")
		defer fmt.Println("老王")
		defer fmt.Println("你好")
		//defer執行順序
		//一個函數有多個defer語句 它們會以後進先出的順序執行
	*/

	/*
		//即使函數或者某個延遲調用發生錯誤 這些調用依舊會被執行
		defer fmt.Println("hello")
		defer fmt.Println("老王")
		defer test87(0)
		defer fmt.Println("你好")
	*/

	a, b := 10, 20
	defer func() {
		fmt.Println("1匿名函數a", a)
		fmt.Println("1匿名函數b", b)
	}()
	defer func(a, b int) {
		fmt.Println("2匿名函數a", a)
		fmt.Println("2匿名函數b", b)
	}(a, b)

	a, b = 100, 200
	fmt.Println("main函數a", a)
	fmt.Println("main函數b", b)
}
