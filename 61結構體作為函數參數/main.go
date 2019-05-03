package main

import (
	"fmt"
)

//Hero結構體
type Hero struct {
	name  string
	age   int
	power int
}

func test61(h Hero) {
	h.power = 120
	fmt.Println(h)
}

func main() {
	h := Hero{"鋼鐵俠", 30, 100}

	//值傳遞
	test61(h)

	fmt.Println(h)
}
