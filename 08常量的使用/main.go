package main

import "fmt"

func main() {

	//變量聲明 var

	//常量聲明 const

	const a float32 = 10
	//a = 20  //err 常量不允許修改
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	//自動推導類型
	const b, c = 3.14, 3.2 // 沒有使用:=
	fmt.Println(b, c)
	fmt.Printf("%T", b)

}
