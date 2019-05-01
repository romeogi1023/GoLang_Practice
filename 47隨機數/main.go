package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main47() {
	//1.導入頭文件  math/rand
	//2.隨機數種子
	//3.創建隨機數

	//創建隨機數種子
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())    //生成比較大隨機數
	fmt.Println(rand.Intn(10)) //常用  取模10 (0-9)
}

//練習 雙色球
func main() {

	//隨機數

	//紅球  1-33  選擇6個  不重覆  藍球  1-16  選擇1個   藍球可以和紅球重覆

	//創建隨機數種子
	rand.Seed(time.Now().UnixNano())

	var red [6]int
	for i := 0; i < len(red); i++ {
		v := rand.Intn(33) + 1

		for j := 0; j < i; j++ {
			//數據重覆
			if v == red[j] {
				//重新隨機
				v = rand.Intn(33) + 1
				//
				j = -1
			}

		}
		red[i] = v
	}
	fmt.Println("紅球", red, "藍球", rand.Intn(16)+1)
}
