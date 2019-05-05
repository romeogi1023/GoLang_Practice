package main

import (
	"bufio"
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

	//createFile()

	//唯讀方式打開文件
	fp, err := os.Open("./xyz.txt")
	if err != nil {
		fmt.Println("打開文件失敗")
		return
	}

	//關閉文件
	defer fp.Close()

	//創建文件的緩衝區
	r := bufio.NewReader(fp)

	// //行讀取 截取的標誌
	// slice, _ := r.ReadBytes('\n')
	// fmt.Println(string(slice))
	// slice, _ = r.ReadBytes('\n')
	// fmt.Println(string(slice))
	// slice, _ = r.ReadBytes('\n')
	// fmt.Println(string(slice))

	// for {
	// 	//遇到'\n'結束讀取 但是'\n'也讀取進入
	// 	buf, err1 := r.ReadBytes('\n')
	// 	//先打印再判斷
	// 	fmt.Println(string(buf))
	// 	if err1 != nil {
	// 		if err1 == io.EOF {
	// 			break
	// 		}
	// 		fmt.Println(err1)
	// 	}
	// }

	for {
		str, err2 := r.ReadString('\n')
		fmt.Println(str)
		if err2 != nil {
			if err2 == io.EOF {
				break
			}
		}
	}

}
