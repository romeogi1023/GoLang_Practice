package main

import "fmt"

func main50() {
	//數組長度固定
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

func main50_01() {
	//切片
	//定義

	//第一種定義方式
	var slice []int //空切片
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	fmt.Println(len(slice)) //長度
	fmt.Println(cap(slice)) //容量

	//第二種定義方式
	var s1 []int = []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Println(len(s1)) //長度
	fmt.Println(cap(s1)) //容量

	//第三種定義方式 自動推導
	s2 := []int{1, 2, 3}
	fmt.Println(s2)
	fmt.Println(len(s2)) //長度
	fmt.Println(cap(s2)) //容量

	//第四種定義方式
	//長度 小於 容量
	//可省略容量
	s := make([]int, 5, 10)
	fmt.Println(s)
	fmt.Println(len(s)) //長度
	fmt.Println(cap(s)) //容量
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(s)
}

func main() {
	//給值

	//下標給值
	s := make([]int, 5, 10)
	s[0] = 1
	s[1] = 2
	fmt.Println(s)

	//循環給值
	s1 := make([]int, 5, 10)
	for i := 0; i < len(s1); i++ {
		s1[i] = i
	}
	fmt.Println(s1)

	for k, v := range s1 {
		fmt.Println(k, v)
	}
}
