package main

import (
	"fmt"
)

type student07 struct {
	name  string
	age   int
	score int
}

//對象不同 方法名相同 不會衝突
//在方法調用中 方法接收者是指針類型
//指針類型 普通類型 表示的是相同對象的類型
func (s *student07) Print() {
	s.score = -9
	fmt.Println(*s)
}

func main() {
	stu := student07{"亞索", 30, -5}
	//值傳遞(s student07)	地址傳遞(s *student07)
	stu.Print()
	fmt.Println(stu)
}
