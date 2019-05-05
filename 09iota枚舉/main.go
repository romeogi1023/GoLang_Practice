package main

import "fmt"

func main() {

	//1.iota常量自動生成器，每個一行，自動加1
	//2.iota給常量賦值使用
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)

	fmt.Println(a, b, c)

	//3.iota遇到const，重置為0
	const d = iota

	fmt.Println(d)
	//4.可以只寫一個iota  在一個括號
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1)

	//5.如果在同一行 ，值一樣
	const (
		i          = iota             //0
		j1, j2, j3 = iota, iota, iota // 1 1 1
		k          = iota             //2

	)
	fmt.Println(i)
	fmt.Println(j1, j2, j3)
	fmt.Println(k)
}
