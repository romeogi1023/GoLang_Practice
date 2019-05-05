package main

import "fmt"

//函數定義   形參
//參數名  自己起   ... type  不定參
func sum(a int, args ...int) {

	//fmt.Println(args)
	//sum:= args[0]+args[1]+args[2]+args[3]
	//fmt.Println(sum)

	//fmt.Println(len("hello"))

	//count:=len(args)
	//fmt.Println(count)

	sum := 0
	//i 下標

	//for循環遍歷
	//for i:=0;i<len(args) ; i++ {
	//	//fmt.Println(args[i])
	//	sum+=args[i]
	//}
	//fmt.Println(sum)

	//匿名變量  丟棄數據
	for _, v := range args {
		//fmt.Println(v)
		sum += v
	}
	fmt.Println(sum, "hello", 444, 55, 5, 5, 5, 5, 5, 5, 5, 5)
}

func main() {

	//函數調用
	sum(1)
}

//1.不定參放最後
//2.函數調用  固定參數必須傳值  不定參數 根據需要決定是否傳值
