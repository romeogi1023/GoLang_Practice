package main

import "fmt"

func test11(args ...int) {
	fmt.Println(args)
	fmt.Println(args[0])
}

func main() {
	test11(1, 2, 3, 4, 5)
}
