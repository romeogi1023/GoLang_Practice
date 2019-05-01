package main

import (
	"fmt"
)

func test36() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	//參數傳遞
	//第一種
	func(a, b int) {
		var sum int
		sum = a + b
		fmt.Println(sum)
	}(3, 6)

	//第二種
	ff := func(a, b int) {
		var sum int
		sum = a + b
		fmt.Println(sum)
	}
	ff(2, 4)

	//匿名函數返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
	fmt.Println(x, y)

	//閉包
	myFunc := test36()
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())

}
