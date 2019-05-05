package main

import "fmt"

func main0101() {
	//數組  是指壹系列同壹類型數據的集合
	//數組定義
	//var 數組名 [長度]類型
	var a [10]int
	//fmt.Println(len(a))
	fmt.Println(a)

	//註意以下方式
	//報錯err
	//var arr int = 10
	//var a [n]int

	//數組賦值
	//a[0] = 1
	//a[1] = 2
	//a[2] = 3
	//a[3] = 4
	//fmt.Println(a)
	//fmt.Println(a[0])
	//fmt.Println(a[1])
	//下標超範圍  數組越界
	//fmt.Println(a[10])
	//for循環 數組賦值 輸出
	for i := 0; i < 10; i++ {
		a[i] = i + 1
	}
	fmt.Println(a)

	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}
	//匿名變量
	for _, v := range a {
		fmt.Println(v)
	}
	//數組存儲的元素類型可以是其他類型
	//定義完成  直接輸出
	//var a [10]float64  //直接輸出  默認為0
	//var a [10]string  //直接輸出  默認為空字符
	//var a [10]bool  //直接輸出  默認為false
}

func main0102() {
	//數組初始化
	//1.全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	//自動推導
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//部分初始化
	//沒有初始化的部分 默認為0
	c := [5]int{1, 2, 3}
	fmt.Println(c)

	//指定某個元素初始化

	d := [5]int{2: 10, 4: 20}
	fmt.Println(d)
	//...  通過初始化確定長度
	f := [...]int{1, 2, 3}
	fmt.Println(len(f))

}

func main() {

	//常見問題
	//數組長度  常量
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	//數組下標  數組下標越界
	//arr[6] = 123  //err
	//arr[-1] = 123 //err
	//數組名
	//arr = 123 //err

	//兩個數組類型相同 個數相同  可以賦值
	//arr1 := arr
	//
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//
	//fmt.Printf("%T\n",arr)

	//數組名表示 整個數組  數組名對應地址 就是數組第1個元素的地址
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[2])

}
