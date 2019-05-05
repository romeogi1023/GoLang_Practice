package main

import "fmt"

func main() {
	//變量的初始化
	//var score int = 700
	//自動推導類型
	score := 699

	//if語法結構

	//if 條件表達式 {    //左括號和if在同壹行
	//	代碼塊
	//}
	//條件表達式沒有括號
	if score == 700 {
		fmt.Println("去上清華北大")
	}

	//if支持1個初始化語句，初始化語句和判斷條件用 ; 分割
	if a := 700; a == 700 {
		fmt.Println("去上清華北大")
	}

}
