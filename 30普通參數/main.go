package main

import "fmt"

//定義函數			//形參
func test10(a int, b string) {
	fmt.Printf("a=%d,b=%s\n", a, b)
}

func main() {

	//函數調用		//參數的傳遞
	test10(1, "H") //  實參
	test10(3, "E")

}
