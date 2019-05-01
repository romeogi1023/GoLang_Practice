package main

import "fmt"

func main() {
	var arr [10]int = [10]int{9, 1, 4, 2, 3, 6, 5, 7, 8, 10}

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
