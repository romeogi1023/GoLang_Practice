package main

import (
	"fmt"

	. "github.com/ahmetb/go-linq"
)

func main() {

	testOrderBy()

	testSingleWith()

}

// testSingleWith 找出長度大於10 的字串
func testSingleWith() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}

	fruit := From(fruits).
		SingleWith(
			func(f interface{}) bool { return len(f.(string)) > 10 },
		)

	fmt.Println(fruit)
}

// testOrderBy 由小到大排序
func testOrderBy() {
	input := []int{10, 20, 40, 15, 18, 11, 25, 33}
	var output []int

	From(input).OrderBy(func(i interface{}) interface{} { return i }).ToSlice(&output)

	fmt.Println(output)
}
