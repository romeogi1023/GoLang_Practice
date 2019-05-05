package main

import "fmt"

func main() {

	//1
	if score := 500; score > 200 {
		fmt.Println("去上大學")
	} else { //else 後面沒有條件
		fmt.Println("去藍翔")
	}

	//2常用   效率高
	a := 10
	if a == 10 {
		fmt.Println("a==10")

	} else if a > 10 {
		fmt.Println("a>10")
	} else {
		fmt.Println("a<10")
	}
	//3.
	b := 10
	if b == 10 {
		fmt.Println("b==10")
	}
	if b > 10 {
		fmt.Println("b>10")
	}
	if b < 10 {
		fmt.Println("b<10")
	}

}
