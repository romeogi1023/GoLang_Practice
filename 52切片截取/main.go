package main

import "fmt"

func main52() {
	//初始化
	s := []int{10, 20, 30, 40, 50}
	fmt.Println(cap(s))
	//截取 s[low:high:max]
	//len = high - low
	//cap = max - low
	slice := s[2:4:5]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5]
	//截取後的切片 還是原始切片中的一塊內容  修改截取後的切片 影響原始切片
	s1[2] = 999
	fmt.Println(s1)
	fmt.Println(s)

	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", s1)
}
