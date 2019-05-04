package main

import (
	"fmt"
)

func test(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		fmt.Println("err=", err)
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := test(10, 0)
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println(result)
	}
}
