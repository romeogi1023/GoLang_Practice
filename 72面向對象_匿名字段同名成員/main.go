package main

import (
	"fmt"
)

//結構體嵌套結構體
//父類
type Person2 struct {
	id   int
	name string
	age  int
}

//子類
type Student2 struct {
	Person2 //匿名字段
	name    string
	score   int
}

func main() {

	var s1 Student2

	//子類賦值
	//子類和父類結構體有相同的成員名 默認賦值給子類  採用就近原則
	s1.name = "花花"

	//父類賦值
	s1.Person2.name = "David"

	fmt.Println(s1)
}
