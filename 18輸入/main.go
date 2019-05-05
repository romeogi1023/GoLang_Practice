package main

import "fmt"

func main() {

	var a float64
	fmt.Printf("請用戶輸入數據：")
	//阻塞 等待用戶的輸入
	//格式化的輸入
	//fmt.Scanf("%f",&a)

	//更簡單

	fmt.Scan(&a)

	fmt.Println(a)

}
