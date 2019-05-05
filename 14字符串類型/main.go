package main

import "fmt"

func main() {

	var str1 string //聲明變量
	str1 = "abc"    //賦值
	fmt.Println(str1)

	//自動推導類型
	str2 := "知了課堂"
	fmt.Println(str2)
	fmt.Printf("%T\n", str2)

	//ch := 'a'
	//str := "a"   //'a''\0'  字符串結束標誌

	//len函數  計算字符串中字符的個數  不包含\0
	//在go語言中 1個漢字占3個字符
	fmt.Println(len(str2))

	//字符串拼接  +

	str3 := "hello"
	str4 := "知了課堂"
	str5 := str4 + str3
	fmt.Println(str5)

}
