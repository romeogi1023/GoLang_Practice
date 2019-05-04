package main

import (
	"fmt"
	"os"
)

func main() {

	//建立文件
	//os.Create(文件名)  文件名  可以寫絕對路徑 和 相對路徑
	//返回值	文件指針 錯誤訊息

	//fp, err := os.Create("xyz.txt")		//建立文件於 工作目錄
	//fp, err := os.Create("D:/xyz.txt")	//建立文件於 絕對路徑
	//fp, err := os.Create("D:\\xyz.txt")	//建立文件於 絕對路徑
	//fp, err := os.Create("/xyz.txt") 		//建立文件於 在d糟
	//fp, err := os.Create("../xyz.txt") 	//建立文件於 工作目錄的上層目錄
	fp, err := os.Create("./xyz.txt") //建立文件於 工作目錄 (相對路徑)

	if err != nil {
		//文件創建失敗
		/*
			1.路徑不存在
			2.文件權限
			3.程序打開文件上限
		*/
		fmt.Println("文件創建失敗")
		return
	}

	//關閉文件
	//如果打開文件不關閉 1.造成內存浪費 2.程序打開文件的上限
	//fp.Close()
	defer fp.Close()
}
