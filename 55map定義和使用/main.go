package main

import (
	"fmt"
)

func main() {
	// names := []string{"張三", "李四", "王五"}
	// fmt.Println(names[2])

	//map[鍵類型]值類型

	//定義1
	var m map[int]string
	fmt.Println(m)
	//內存地址編號為0的空間 系統占用  不允許進行讀寫操作
	fmt.Printf("%p\n", m)
	//在字典中不能使用cap函數 只能使用len函數
	//fmt.Println(cap(m))
	fmt.Println(len(m))

	//定義2
	// m2 := make(map[int]string, 3)
	m2 := make(map[int]string)
	fmt.Println(m2)
	fmt.Printf("%p\n", m2)

	//賦值
	//鍵是值一的 值可以重覆
	//map是無序的 所以打印出來的順序 不一定會照順序
	//map長度是自動擴容
	m2[1] = "張三"
	fmt.Println(len(m2))
	m2[2] = "李四"
	fmt.Println(len(m2))
	m2[3] = "王五"
	fmt.Println(len(m2))
	m2[4] = "張三"
	fmt.Println(len(m2))

	fmt.Println(m2)

	//初始化
	m3 := map[int]string{1: "張三", 2: "李四", 3: "王五"}
	fmt.Println(m3[3])
}
