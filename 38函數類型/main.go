package main

import "fmt"

func demo1(a, b int) {
	fmt.Println(a + b)
}

//定義函數類型  為已存在的數據類型起別名
type FUNCDEMO func(int, int)

func main() {
	//demo1(10,20)
	//函數名表示壹個地址 函數在代碼區的地址

	fmt.Println(demo1)
	//自動推導類型
	//f:=demo1

	//f是func(int,int)函數類型定義的變量
	//var f func(int ,int)
	var f FUNCDEMO
	f = demo1
	fmt.Println(f)
	//通過f調用了函數
	f(10, 20) //  相當於 deme1(10,20)
	//f類型
	fmt.Printf("%T\n", f)

}
