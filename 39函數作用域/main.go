package main

import "fmt"

//全局變量
//1.定義在函數外部的變量  全局變量
//	全局變量作用域 整個文件
//2.全局變量名不能其他文件中的額全局變量名重名
//全局變量名和局部變量名可以重名  盡量避免

var a int = 20

func test8() {
	a = 5
	a += 1
}

func demo6() {
	fmt.Println(a)
}

func main02() {

	//作用域  作用範圍
	//局部變量
	//1.定義在函數內部的變量  局部變量
	//	局部變量的作用域  函數內部有效
	//2.不同的函數  可以定義相同的局部變量名  互不影響
	//3.變量 先聲明  在使用  在函數內部  變量名是唯壹的
	//var i int = 10
	//for i:=0;i<5 ;i++  {
	//	fmt.Println(i)
	//}
	//fmt.Println(i)

	//如果全局變量名和局部變量名相同 那麽會使用 局部變量 就近原則
	a := 9
	//test8()
	fmt.Println(a)

	//修改全局變量  會影響到其他位置使用全局變量
	demo6()

}

func main() {
	//打印函數地址       代碼區
	fmt.Println(demo6)
	//打印全局變量地址    數據區
	fmt.Println(&a)
	//打印局部變量地址   棧區
	a := 10
	fmt.Println(&a)
}
