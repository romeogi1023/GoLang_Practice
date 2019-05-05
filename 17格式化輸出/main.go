package main

import (
	"fmt"
)

func main() {

	a := 10
	b := "abc"
	c := 'a'
	d := 3.1415
	//%T 打印變量所屬數據類型

	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d)

	//%d 整型格式
	//%s 字符串格式
	//%c 字符格式
	//%f 浮點型格式

	fmt.Printf("%d,%s,%c,%.2f\n", a, b, c, d)

	//%v自動格式匹配輸出

	fmt.Printf("%v,%v,%v,%v", a, b, c, d)

}
