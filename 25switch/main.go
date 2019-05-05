package main

import "fmt"

func main() {

	//語法
	//switch 變量(表達式) {   和switch關鍵字同行
	//case 值1:
	//	代碼體
	//case 值2:
	//	代碼體
	//default:
	//	代碼體
	//}
	//1.支持多個條件的匹配
	//2.不同的case之間不需要break
	var score int = 90
	switch score {

	case 90:
		fmt.Println("A")
		//fallthrough
	case 80:
		fmt.Println("B")

	case 50, 60, 70:
		fmt.Println("C")

	default:
		fmt.Println("D")
	}

	//1.

	switch s1 := 90; s1 { //初始化語句;條件
	case 90:
		fmt.Println("A")
		//fallthrough
	case 80:
		fmt.Println("B")

	case 50, 60, 70:
		fmt.Println("C")

	default:
		fmt.Println("D")
	}

	var s2 int = 90
	switch { //沒有條件
	case s2 >= 90: //寫判斷語句
		fmt.Println("A")
	case s2 >= 80:
		fmt.Println("B")
	default:

		fmt.Println("C")

	}

	//3.

	switch s3 := 90; { //只有初始化語句  沒有條件
	case s3 >= 90: //寫判斷語句
		fmt.Println("A")
	case s3 >= 80:
		fmt.Println("B")
	default:

		fmt.Println("C")
	}
}

//總結

/*
優點
1.if 適合進行區間的判斷  嵌套使用
2.switch  固定值    執行效率高  可以將多個滿足相同條件的值放在壹起

缺點：
1.if 執行效率低
2.switch 不建議使用嵌套
*/
