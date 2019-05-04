package main

import (
	"fmt"
	"strconv"
)

func main9101() {
	str := "hello world"
	//將字符串轉成字符切片	強制類型轉換
	slice := []byte(str)
	fmt.Println(slice)
}

func main9102() {
	//字符切片
	slice := []byte{'h', 'e', 'l', 'l', 'o', 97}
	fmt.Println(slice)
	fmt.Println(string(slice))
}

func main9103() {

	//將其它類型轉成字符串	Format

	fmt.Println(" \n=== 將bool轉字串 === ")
	strbool := strconv.FormatBool(true)
	fmt.Println(strbool)
	fmt.Printf("%T\n", strbool)

	fmt.Println(" \n=== 將int轉字串 === ")
	strint := strconv.FormatInt(120, 10) //計算機中進制  可以表示2-36進制  一般常用 2,8,10,16進制
	fmt.Println(strint)
	fmt.Printf("%T\n", strint)

	fmt.Println(" \n=== 將float轉字串 === ")
	strfloat := strconv.FormatFloat(3.14159, 'f', 4, 64) //strconv.FormatFloat(浮點小數值, 顯示方式, 顯示小數點後幾位同時會四捨五入, 最大小數位數float64)
	fmt.Println(strfloat)
	fmt.Printf("%T\n", strfloat)

	fmt.Println(" \n=== 將itoa轉字串 === ")
	stritoa := strconv.Itoa(123)
	fmt.Println(stritoa)
	fmt.Printf("%T\n", stritoa)
}

func main() {

	//字符串轉成其它類型
	/*
		b, err := strconv.ParseBool("true")
		if err != nil {
			fmt.Println("類型轉換出錯")
		} else {
			fmt.Println(b)
			fmt.Printf("%T\n", b)
		}
	*/

	/*
		//v, err := strconv.ParseInt("abc", 10, 64)
		v, err := strconv.ParseInt("abc", 16, 64)
		fmt.Println(v, err)
	*/

	/*
		v, err := strconv.ParseFloat("a3.14159", 64)
		fmt.Println(v, err)
	*/

	v, _ := strconv.Atoi("123aa")
	fmt.Println(v)

}
