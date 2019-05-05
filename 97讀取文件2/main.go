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

	//唯讀方式打開文件
	fp, err := os.Open("./xyz.txt")
	if err != nil {
		fmt.Println("打開文件失敗")
		return
	}

	buf := make([]byte, 1024*2) //2k

	for {
		n, err := fp.Read(buf)
		//io.EOF文件的結尾
		//讀取到文件末尾 返回值 errors.New("EOF")
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}

	//關閉文件
	defer fp.Close()
}
