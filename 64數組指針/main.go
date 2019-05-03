package main

import (
	"fmt"
)

func main6401() {

	arr := [3]int{1, 2, 3}
	//定義指針  指向數組  數組指針
	var p *[3]int
	//指針和數組建立關係
	p = &arr
	fmt.Println(*p)

	//自動推導類型創建數組指針
	//p := &arr

	//通過指針操作數組

	// (*數組指針變量)[下標] = 值
	(*p)[0] = 111
	// 直接操作
	p[1] = 222

	fmt.Println(*p)
}

func main6402() {
	arr := [5]int{1, 2, 3, 4, 5}
	//指針變量和要存儲數據類型要相同
	p1 := &arr
	p2 := &arr[0]
	//p1和p2在內存中指向相同的地址
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", p2)

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
}

func swaps(p *[5]int) { //指針數組
	(*p)[0] = 111
	//p[0] = 111
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	swaps(&a) //傳地址

	fmt.Println(a)

}
