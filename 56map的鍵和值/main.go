package main

import (
	"fmt"
)

func main56() {

	//map中key類型必須支持 == != 一般建議寫 基本類型
	//切片 函數 切片的結構類型  不能作為字典的key
	//m := make(map[[]string]int) //invalid map key type []string
	//map中 數據是無序的
	m := make(map[string][3]int)
	m["小明"] = [3]int{100, 99, 98}
	m["小亮"] = [3]int{6, 9, 8}
	m["小紅"] = [3]int{101, 100, 102}

	fmt.Println(m)
	fmt.Println(m)

	//for循環取值
	for k, v := range m {
		fmt.Println(k, v)
	}

}

func main() {

	m := make(map[int]string)
	m[1] = "劉備"
	m[2] = "曹操"
	m[3] = "孫權"

	//通過key取值
	fmt.Println(m[1])
	//fmt.Println(m["劉備"]) //cannot convert "劉備" (type untyped string) to type int

	//在map中如果沒有提供key找到具體的值 打印value類型 的默認值
	fmt.Println(m[4])

	//判斷鍵值對是否存在
	value, ok := m[2]
	fmt.Println(value, ok)

	//刪除 map中的元素  根據key進行刪除
	delete(m, 2)
	//如果key不存在 不會報錯
	delete(m, 4)
	fmt.Println(m)
}
