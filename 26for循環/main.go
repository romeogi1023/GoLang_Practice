package main

import "fmt"

func main() {

	// 計算 1+2+3+...+10的和
	sum := 0
	for i := 0; i < 11; i++ {
		sum += i
		//fmt.Println(sum)
	}
	fmt.Println("加總和 = ", sum)

	//通過for輸出
	str := "abc"
	for x := 0; x < len(str); x++ {
		fmt.Println(str[x])
		fmt.Printf("str[%d]=%c\n", x, str[x])
	}

	//通過range輸出(如同foreach)
	str2 := "abc"
	for n, data := range str2 {
		fmt.Println(n, data)
	}
	for _, data := range str2 {
		fmt.Println(data)
	}
}
