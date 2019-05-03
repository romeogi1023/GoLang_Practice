package main

import (
	"fmt"
)

func main() {

	/*
		var i int
		i = 100
		fmt.Printf("%d\n", i)
		fmt.Printf("%p\n", &i)
		//fmt.Printf("%v\n", &i)

		//指針變量 定義

		var p *int //定義指針變量
		p = &i     //把變量i 地址複製指針變量p
		fmt.Printf("i=%d,p=%p\n", i, p)

		*p = 80 //使用指針修改值
		fmt.Printf("i=%d,p=%p\n", i, p)
		fmt.Printf("i=%d,p=%v\n", i, *p)
	*/

	//注意
	//1.默認值 nil

	//內存地址編號為0  0-255 空間  系統占用  不允許用戶訪問

	// //空指針
	// var p *int = nil
	// fmt.Println(p)

	// //野指針  指針變量指向一個未知的空間   會報錯
	// //程序不允許出現野指針
	// var a int
	// var p *int
	// p = &a //沒有指向  直接操作
	// *p = 56
	// fmt.Println(p, *p)

	//new函數
	//gc垃圾回收機制  自動釋放空間
	var p *int
	p = new(int)
	*p = 57
	fmt.Println(*p)

	//自動推導類型
	q := new(int)
	*q = 999
	fmt.Println(*q)

}
