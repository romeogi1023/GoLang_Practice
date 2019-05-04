package main

import (
	"fmt"
)

type person10 struct {
	id   int
	name string
	age  int
}

func (p person10) PrintInfo1() {
	fmt.Printf("info1: %p, %v\n", &p, p)
}

//建議使用 指針類型
func (p *person10) PrintInfo2() {
	fmt.Printf("info2: %p, %v\n", p, *p)
}

func main() {
	p := person10{1, "make", 22}
	// p.PrintInfo1() //0xc000054440, {1 make 22}
	// p.PrintInfo2() //0xc000054420, {1 make 22}

	// fmt.Println(p.PrintInfo1)
	// fmt.Println(p.PrintInfo2)
	// fmt.Printf("%T\n", p.PrintInfo1)
	// fmt.Printf("%T\n", p.PrintInfo2)

	// //方法值  隱式傳遞  隱藏的是接收者  綁定實例(對象)
	// var pfunc1 func()
	// pfunc1 = p.PrintInfo1
	// pfunc1()
	// fmt.Printf("%T\n", pfunc1)

	// //方法表達式  顯式傳參
	// pfunc1 := person10.PrintInfo1
	// pfunc1(p)
	// pfunc2 := (*person10).PrintInfo2
	// pfunc2(&p)
	// fmt.Printf("%T\n", pfunc1)
	// fmt.Printf("%T\n", pfunc2)

	//對象相同 但是函數類型不同  不能賦值
	//函數類型相同  可以賦值
	pfunc1 := p.PrintInfo1
	pfunc1 = p.PrintInfo2
	pfunc1()
	fmt.Printf("%T\n", pfunc1)
}
