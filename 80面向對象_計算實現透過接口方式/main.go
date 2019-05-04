package main

import (
	"fmt"
)

//定義接口
type Opter interface {
	//方法聲明
	Result() int
}

//父類
type Operate struct {
	num1 int
	num2 int
}

//加法子類
type Add struct {
	Operate
}

//加法子類的方法
func (a *Add) Result() int {
	return a.num1 + a.num2
}

//減法子類
type Sub struct {
	Operate
}

//減法子類的方法
func (s *Sub) Result() int {
	return s.num1 - s.num2
}

//多態的實現
func Result(o Opter) {
	v := o.Result()
	fmt.Println(v)
}

func main() {

	/*
		//1.通過對象方法調用
		//創建加法
		var a Add
		a.num1 = 10
		a.num2 = 20
		v := a.Result()
		fmt.Println(v)

		//創建減法
		var s Sub
		s.num1 = 10
		s.num2 = 20
		r := s.Result()
		fmt.Println(r)
	*/

	/*
		//2.通過接口實現
		var o Opeter
		var a Add = Add{Operate{10, 20}}
		o = &a
		value := o.Result()

		fmt.Println(value)
	*/

	//3.多態實現
	var a Add = Add{Operate{10, 20}}
	Result(&a)

	var s Sub = Sub{Operate{10, 20}}
	Result(&s)
}
