package main

import "fmt"

func main() {

	//a := 10
	//
	//fmt.Println(a)
	//a--
	//a--
	//fmt.Println(a)
	//b := 20
	//fmt.Println(a == b)
	//fmt.Println(true || false)
	////fmt.Println(!true)

	//var a int= 60

	//a := 60    //  0011 1100
	//b := 13	   //  0000 1101
	//var c int
	//
	//c = a & b
	//fmt.Println(c)  //  0000 1100
	//
	//fmt.Printf("%p\n",&a)
	//fmt.Println(a)
	//fmt.Printf("%d\n",*&a)

	a := 20
	b := 10
	c := 15
	d := 5

	var e int

	e = (a + b) * c / d
	fmt.Println(e)
	e = ((a + b) * c) / d

}
