package main

import (
	"fmt"
	"io"
	"os"
)

func main9401() {

	fp, err := os.Create("./xyz.txt")
	//建立文件
	if err != nil {
		fmt.Println("文件創建失敗")
		return
	}

	//寫文件
	// \n在windows中可能不會換行  原因 在windows文本文件中換行\r\n 回車     在linux中換行\n
	fp.WriteString("hello world\n")
	fp.WriteString("性感荷官在線發牌")

	//關閉文件
	//如果打開文件不關閉 1.造成內存浪費 2.程序打開文件的上限
	//fp.Close()
	defer fp.Close()
}

func main9402() {

	fp, err := os.Create("./xyz.txt")
	if err != nil {
		fmt.Println("文件創建失敗")
		return
	}

	//寫操作
	// slice := []byte{'h', 'e', 'l', 'l', 'o'}
	// count, err1 := fp.Write(slice)
	count, err1 := fp.Write([]byte("性感老王在線授課"))
	if err1 != nil {
		fmt.Println("寫入文件失敗")
		return
	} else {
		fmt.Println(count)
	}

	defer fp.Close()

}

func main() {

	fp, err := os.Create("./xyz.txt")
	if err != nil {
		fmt.Println("文件創建失敗")
		return
	}

	//寫操作

	//獲取光標流位置
	//獲取文件起始到結尾有多少個字符
	//count, _ := fp.Seek(0, os.SEEK_END)	//os.SEEK_END 未來可能不再支持
	count, _ := fp.Seek(0, io.SeekEnd)
	fmt.Println(count)

	fp.WriteAt([]byte("hello world"), count)
	fp.WriteAt([]byte("hahaha"), 0)
	fp.WriteAt([]byte("秀兒"), 19)

	defer fp.Close()

}
