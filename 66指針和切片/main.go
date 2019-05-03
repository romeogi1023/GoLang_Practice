package main

import (
	"fmt"
)

func main66() {

	s := []int{1, 2, 3, 4, 5}
	//指針和切片建立聯系
	p := &s
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", s)

	//*[]int
	//var p *[]int
	fmt.Printf("%T\n", p)

	//通過指針操作切片元素
	//p[1] = 222  //err
	(*p)[1] = 222
	fmt.Println(s)

	//for循環遍歷
	for i := 0; i < len(s); i++ {
		fmt.Println((*p)[i])
	}
}

//切片指針作為函數參數
func test66(s *[]int) {
	*s = append(*s, 6, 7, 8, 9)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	//地址傳遞
	test66(&s)
	fmt.Println(s)
}
