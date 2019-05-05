package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var srcFileName string
	var dstFileName string
	fmt.Printf("請輸入源文件名稱：")
	fmt.Scan(&srcFileName)
	fmt.Printf("請輸入目的文件名稱：")
	fmt.Scan(&dstFileName)

	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件不能同名")
		return
	}

	//唯讀方式打開源文件
	sf, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("打開源文件失敗", err1)
		return
	}

	//新建目的文件
	df, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("新建目的文件失敗", err2)
		return
	}

	//操作完畢 關閉文件
	defer sf.Close()
	defer df.Close()

	//核心處理 從源文件讀取內容 往目的文件寫入 讀取多少 寫入多少
	buf := make([]byte, 1024*4) //4k

	for {
		n, err := sf.Read(buf)
		if err != nil {

			if err == io.EOF { //文件讀取完畢
				break
			}
		}
		//寫入
		df.Write(buf[:n])
	}

}
