package main

import (
	"fmt"
)

type hero struct {
	name  string
	age   int
	power int
}

func test60(m map[int]hero) {

	// //err
	// m[102].power = 89

	stu := m[102]
	stu.power = 89

	m[102] = stu

	fmt.Println(m)
	fmt.Printf("%p\n", m)
}

func main6001() {

	//將結構體作為map中的值 value
	m := make(map[int]hero)

	//map中的數據不建議排序操作
	m[101] = hero{"鋼鐵俠", 30, 100}
	m[102] = hero{"美國隊長", 30, 90}

	fmt.Println(m)

	delete(m, 102)
	fmt.Println(m)
}

func main6002() {

	//value 類型 是一個切片
	m := make(map[int][]hero)

	m[101] = []hero{{"鋼鐵俠", 30, 100},
		{"蜘蛛俠", 17, 80}}
	m[101] = append(m[101], hero{"星爵", 30, 10})
	fmt.Println(m)

	m[102] = []hero{{"美國隊長", 30, 90}}
	m[102] = append(m[102], hero{"冬兵", 30, 75})
	fmt.Println(m)
	fmt.Println(m[102][1])
}

func main() {

	m := make(map[int]hero)
	//map中的數據不建議排序操作
	m[101] = hero{"鋼鐵俠", 30, 100}
	m[102] = hero{"美國隊長", 30, 90}

	//將map作為函數參數  地址傳遞
	test60(m)

	fmt.Println(m)
	fmt.Printf("%p\n", m)
}
