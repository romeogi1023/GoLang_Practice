package main

import "fmt"

func main82() {

	var i interface{}
	//接口類型可以接收任意類型的數據

	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = 10
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = 3.14
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = "Go"
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	fmt.Println(10, 3.14, "Go", 'a')
}

func test82() {
	fmt.Println("test")
}

func main() {

	//空接口類型的切片
	var i []interface{}

	fmt.Printf("%T\n", i)

	i = append(i, 10, 3.14, "aaa", test82)

	fmt.Println(i)

	for idx := 0; idx < len(i); idx++ {
		fmt.Println(i[idx])
	}
}
