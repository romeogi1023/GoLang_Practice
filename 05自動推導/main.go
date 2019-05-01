package main

import (
	"fmt"
)

func main() {
	//自動推導
	var a = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	//常用的自動推導
	//自動推導  := 左邊變量名沒使用過
	b := 10
	b = 20
	b = 30
	fmt.Println(b)

	c := 3.14
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	d, f, e := 20, 3.14, 30
	fmt.Println(d, f, e)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", e)
}
