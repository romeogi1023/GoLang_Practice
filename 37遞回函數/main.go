package main

import "fmt"

//遞迴函數 調用函數自己本身
//遞迴函數相同的結構
//1.跳出條件
func test37(a int) {
	if a == 1 {
		fmt.Println(a)
		return //函數結束
	}
	test37(a - 1)
	fmt.Println(a)
}

func sum1(num int) int {
	if num == 1 {
		return 1
	}
	return num + sum1(num-1)
}

func sum2(num int) int {
	if num == 100 {
		return 100
	}
	return num + sum2(num+1)
}

func main() {
	test37(5)
	fmt.Println(sum1(100))
	fmt.Println(sum2(1))
}
