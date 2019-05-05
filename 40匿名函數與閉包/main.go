package main

import "fmt"

func test11() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
func main() {

	//匿名函數  沒有名字的函數
	//var num int
	//num = 9
	//f:=func(){
	//	num++
	//	fmt.Println(num)
	//} //花括號  {}() 函數調用
	//////函數調用
	//f()
	////fmt.Println(num)
	//
	//type FuncType func()
	//
	//var f1 FuncType
	//f1 = f
	//f1()
	//fmt.Println(num)

	//匿名函數參數傳遞
	//1.
	//func(a,b int){  //匿名函數帶參數
	//	var sum int
	//	sum = a+b
	//	fmt.Println(sum)
	//}(3,6)  //調用時傳值
	//2.
	//f:=func(a,b int){  //匿名函數帶參數
	//	var sum int
	//	sum = a+b
	//	fmt.Println(sum)
	//} //調用時傳值
	//f(3,6)

	//匿名函數返回值

	//x,y := func(i,j int)(max,min int){
	//	if i >j {
	//		max = i
	//		min = j
	//	}else {
	//		max = j
	//		min = i
	//	}
	//	return
	//}(10,20)
	//fmt.Println(x,y)

	//閉包
	f := test11()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
