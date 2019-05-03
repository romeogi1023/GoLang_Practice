package main

import (
	"fmt"
)

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
