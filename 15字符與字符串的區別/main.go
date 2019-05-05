package main

import "fmt"

func main() {

	var ch byte
	var str string

	//字符
	//1.單引號
	//2.字符，往往都是只有1個字符  除了 \n \t 轉義字符

	ch = 'a'
	fmt.Println(ch)
	//字符串
	//1.雙引號
	//2.字符串可以有1個或多個字符組成
	//3.字符串都是隱藏了1個結束符  '\0'
	str = "a"
	fmt.Println(str)

	str = "hello go"
	fmt.Println(str[0], str[1])
	fmt.Printf("%c,%c\n", str[0], str[1])

}
