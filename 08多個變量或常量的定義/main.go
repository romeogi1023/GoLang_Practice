package main

import "fmt"

func main() {

	//變量 程序可運行期間 可以改變的量  聲明關鍵字var
	//var a,b int = 10,20
	//fmt.Println(a,b)

	//快捷鍵  ctrl+/ 註釋 解註釋

	//不同類型變量的定義
	//var a int =1
	//var b float64 = 2.0
	//var (
	//	a int =1
	//	b float64 = 2.0
	//)

	//自動推導類型
	//var (
	//	a = 1
	//	b = 2.0
	//)

	a, b := 1, 2.0
	fmt.Println(a, b)

	//常量 ：程序運行期間 不可以更改量   聲明關鍵字 const
	//常量的定義

	//const i int = 10
	//const j float64 = 3.14
	//fmt.Println(i,j)

	//const (
	//	i int = 10
	//	j float64 = 3.14
	//)

	//自動推導

	const (
		i = 10
		j = 3.14
	)
	fmt.Println(i, j)
	fmt.Printf("%T", j)

}
