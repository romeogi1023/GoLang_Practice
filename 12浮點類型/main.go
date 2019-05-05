package main

import "fmt"

func main() {
	//浮點型數據分為 單精度float32 和 雙精度float64
	var a float64 = 123.45655555
	fmt.Println(a)
	var b float32 = 123.45655555
	fmt.Println(b)
	//通過自動推導類型創建的浮點型變量默認為float64
	c := 3.14
	fmt.Println(c)
	fmt.Printf("%T", c)

}
