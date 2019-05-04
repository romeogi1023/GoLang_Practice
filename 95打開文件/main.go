package main

import (
	"fmt"
	"io"
	"os"
)

func createFile() {
	fp, err := os.Create("./xyz.txt")
	if err != nil {
		fmt.Println("文件創建失敗")
		return
	}

	count, _ := fp.Seek(0, io.SeekEnd)
	fmt.Println(count)

	fp.WriteAt([]byte("hello world"), count)
	fp.WriteAt([]byte("hahaha"), 0)
	fp.WriteAt([]byte("秀兒"), 19)

	defer fp.Close()
}

func main() {

	createFile()

	//os.Open 唯讀方式打開
	//fp, err := os.Open("./xyz.txt")

	//os.OpenFile(文件名, 打開方式, 打開權限)
	fp, err := os.OpenFile("./xyz.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("打開文件失敗")
		return
	}

	fp.WriteString("DavidTest")
	fp.WriteAt([]byte("DavidTest"), 40)

	defer fp.Close()

}
