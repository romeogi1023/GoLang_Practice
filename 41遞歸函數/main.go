package main

import "fmt"

//遞歸函數  調用函數自己本身
//遞歸函數相同的結構
//1.跳出條件
func test01(a int) {
	if a == 1 {
		fmt.Println(a)
		return //函數結束
	}
	test01(a - 1)
	fmt.Println(a)
}

func test02(num int) int {
	if num == 1 {
		return 1
	}
	return num + test02(num-1)
}

func test03(num int) int {
	if num == 100 {
		return 100
	}
	return num + test03(num+1)
}

func main() {

	//test01(3)
	fmt.Println(test02(100))
	fmt.Println(test03(1))

}
