package main

import "fmt"

func main53() {
	//append()
	//copy()

	var s []int = []int{1, 2, 3, 4, 5}

	//s1 := make([]int, 5)
	s1 := []int{6, 7, 8, 9}

	fmt.Println(s1)
	copy(s1, s)
	fmt.Println(s1)
	//使用copy進行拷貝 在內存中存儲兩個獨立的切片內容  如果任一個發生修改 不會影響到另一個
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", s1)

	s1[2] = 123
	fmt.Println(s)
	fmt.Println(s1)

	s = append(s, 6, 7, 8, 9)
	fmt.Println(s)
	fmt.Println(s1)
}

func main() {
	var arr []int = []int{8, 7, 5, 6, 4, 2, 3, 1, 9}

	fmt.Println(arr)

	//外層控制
	for i := 0; i < len(arr)-1; i++ {
		//內層控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//滿足條件 進行交換 (大於 升序)  (小於 降序)
			if arr[j] < arr[j+1] {
				//交換數據
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
