package main

import (
	"fmt"
	"strings"
)

func main8901() {
	//查找一個字符串在另一個字符串中是否出現
	str1 := "hello world"
	str2 := "w"

	//Contains(被查找的字符串, 查找的字符串)  返回值 bool
	b := strings.Contains(str1, str2)
	if b {
		fmt.Println("找到了")
	} else {
		fmt.Println("沒有找到")
	}
}

func main8902() {
	//字符串切片
	slice := []string{"123", "456", "789"}
	//Join 字符串的連接
	str := strings.Join(slice, ",")
	fmt.Println(str)
}

func main8903() {
	str1 := "hello world"
	str2 := "l"
	//查找一個字符串在另一個字符串中第一次出現的位置  返回值 int 下標	-1 找不到
	i := strings.Index(str1, str2)
	fmt.Println(i)
}

func main8904() {
	str := "性感網友，在線取名。"
	//將一個字符串重覆n次
	str1 := strings.Repeat(str, 5)
	fmt.Println(str1)
}

func main8905() {
	str := "性感網友在線取名性感性感性感性感性感"
	//字符串替換  屏蔽敏感詞彙
	//如果替換次數小於0	表示全部替換
	// str1 := strings.Replace(str, "性感", "**", 2)
	str1 := strings.Replace(str, "性感", "**", -1)
	fmt.Println(str1)
}

func main8906() {
	str1 := "1300-188-1999"
	//將一個字符串按照標誌位進行切割變成切片
	slice := strings.Split(str1, "-")
	fmt.Println(slice)

	aaa := strings.Join(slice, "xyz")
	fmt.Println(aaa)

	for _, v := range slice {
		fmt.Println(v)
	}
}

func main8907() {
	str := "====a===u=ok===="
	//去掉字符串頭尾的內容
	str1 := strings.Trim(str, "=")
	fmt.Println(str1)
}

func main() {
	str := "    are you ok     "
	//去掉字符串中的空格  並轉成切片  一般用於統計單詞個數
	slice := strings.Fields(str)
	fmt.Println(slice)

	for _, v := range slice {
		fmt.Println(v)
	}
}

/*
總結：
	【查找】
		1. bool類型 := strings.Contains(被查找的字符串, 查找的字符串)
		2. int類型 := strings.Index(被查找的字符串, 查找的字符串)

	【分割】
		[]string類型 := strings.Split(切割字符串, 標誌)

	【組合】
		string類型 := strings.Join(字符串切片, 標誌)

	【重覆】
		string類型 := strings.Repeat(字符串, 次數)

	【替換】
		string類型 := strings.Replace(字符串, 被替換字符串, 替換字符串, 次數)

	【去掉內容】
		1. string類型 := strings.Trim(字符串, 去掉字符串)
		2. []string類型 := strings.Fields(字符串)
*/
