package main

import "fmt"

//切片作為函數參數  返回值
func demo(s []int) []int {
	//fmt.Println(s)
	//s[0] = 123
	//fmt.Println(s)

	//如果函數中使用append進行數據添加，形參地址發生改變，就不會指向實參的地址
	s = append(s, 6, 7, 8, 9)
	fmt.Println(s)
	return s
}

func main() {
	//切片名本身就是個地址
	slice := []int{1, 2, 3, 4, 5}
	//地址傳遞
	slice = demo(slice)
	fmt.Println(slice)
}
