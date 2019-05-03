package main

import "fmt"

func main6901() {
	a := 10
	b := 20

	//一級指針
	p := &a
	fmt.Println(*p)
	// fmt.Println(p)
	// fmt.Printf("%T\n", p)

	//二級指針  指向一級指針的地址
	var pp **int
	pp = &p
	//fmt.Printf("%T\n", pp)

	//通過二級指針間接修改一級指針的值
	*pp = &b
	fmt.Println(*p)

	//通過二級指針間接修改變量的值
	**pp = 100
	fmt.Println(*p)
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	a := 10
	var p *int = &a
	var pp **int = &p
	var ppp ***int = &pp //引用運算符 不能連續使用&&

	fmt.Println(ppp)

	//***ppp = **pp = *p = a

}
