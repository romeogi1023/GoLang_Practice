package main

import "fmt"

func swap(a int, b int) {
	a, b = b, a
}

func main0501() {
	a := 10
	b := 20
	swap(a, b)
	//值傳遞
	fmt.Println(a)
	fmt.Println(b)
}

//數組作為函數參數  返回值
func BubbleSort(arr [10]int) [10]int {
	//外層控制行
	for i := 0; i < len(arr)-1; i++ {
		//內層控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//滿足條件 進行交換  大於 升序  小於  降序
			if arr[j] < arr[j+1] {
				//交換數據
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}
	fmt.Println(arr)
	return arr
}

func main() {

	arr := [10]int{9, 1, 5, 6, 8, 4, 7, 8, 10, 3}
	//數組作為函數參數是值傳遞
	//形參和實參 是不同地址

	//內存中有兩份獨立的數組  儲存不同的數據
	//在函數調用結束後  內存回收  不會影響主函數實參的值

	//如果想通過函數計算結果改變實參的值  需要使用數組作為函數的返回值
	arr1 := BubbleSort(arr)

	fmt.Println(arr)
	fmt.Println(arr1)
}
