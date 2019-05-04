package main

import (
	"fmt"
)

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

func main() {
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
}
