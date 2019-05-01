package main

import (
	"fmt"
)

type Student struct {
	id    int
	name  string
	score int
}

//結構體數組作為函數參數 值傳遞

//結構體切片作為函數參數 地址傳遞
func Sort(arr []Student) {
	//外層控制
	for i := 0; i < len(arr)-1; i++ {
		//內層控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//滿足條件 進行交換 (大於 升序)  (小於 降序)
			if arr[j].score < arr[j+1].score {
				//交換數據
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}

func main() {
	var arr []Student = []Student{
		Student{1, "李白", 100},
		Student{2, "王維", 100},
		Student{3, "杜甫", 100}}

	//fmt.Println(arr)

	// //循環打印結構體信息
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	//修改結構體成員指定信息
	arr[1].score = 97
	arr[2].score = 99

	//結構體切片 添加數據
	arr = append(arr, Student{4, "袁華", 100})

	// //結構體數據作為函數參數 值傳遞
	// //結構體切片作為函數參數 地址傳遞 (引用傳遞)
	Sort(arr)
	fmt.Println(arr)
}
