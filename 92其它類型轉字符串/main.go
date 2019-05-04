package main

import (
	"fmt"
	"strconv"
)

func main() {

	slice := make([]byte, 0, 1024)
	//將其它類型轉成字符串 添加到字符切片裡

	slice = strconv.AppendBool(slice, false)
	slice = strconv.AppendInt(slice, 123, 2)
	slice = strconv.AppendFloat(slice, 3.14159, 'f', 4, 64)
	slice = strconv.AppendQuote(slice, "hello")

	fmt.Println(slice)

	fmt.Println(string(slice))

}
