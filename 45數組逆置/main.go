package main

import "fmt"

//func main() {
//	//練習：從以下整數數組中取出 最大整數、最小整數、總和、平均值。
//	//var a [6]int = [6]int{0,1,2,3,4,5}
//
//	//1.全部初始化
//	var a [5]int = [5]int{1, 2, 3, 4, 5}
//
//	//2.聲明變量
//	var max int = a[0]
//	var min int = a[0]
//	var sum int = a[0]
//
//	//3.循環
//	for i := 0; i < len(a); i++ {
//		if a[i] > max {
//			max = a[i]
//		}
//		if a[i] < min {
//			min = a[i]
//		}
//		sum += a[i]
//	}
//	fmt.Printf("最大值：%d, 最小值：%d, 總和：%d, 平均值：%d", max, min, sum, sum/len(a))
//}

func main() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//數組逆置

	fmt.Println(arr)

	i := 0            //最小下標
	j := len(arr) - 1 //最大下標

	//for 表達式1;表達式2;表達式3{}
	//for 返回值 := range 集合
	//for 條件{}
	for {
		if i >= j {
			//跳出循環
			break
		}
		//交換數組
		arr[i], arr[j] = arr[j], arr[i]
		//改變下標
		i++
		j--
	}

	fmt.Println(arr)
}
