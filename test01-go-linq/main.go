package main

import (
	"fmt"

	. "github.com/ahmetb/go-linq"
)

func main() {

	input := []int{10, 20, 40, 15, 18, 11, 25, 33}
	var output []int

	// 透過 go-linq 排序
	From(input).OrderBy(func(i interface{}) interface{} { return i }).ToSlice(&output)

	fmt.Println(output)
}
