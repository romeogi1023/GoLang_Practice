package main

import "fmt"

func add(s1, s2 int) int {
	sum := s1 + s2
	return sum
}

func main() {
	a, b := 10, 20

	result := add(a, b)
	fmt.Println(result)
}
