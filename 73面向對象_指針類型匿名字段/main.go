package main

import (
	"fmt"
)

//結構體嵌套結構體
//父類
type Person3 struct {
	id   int
	name string
	age  int
}

//子類
type Student3 struct {
	*Person3 //指針類型匿名字段
	score    int
}

func main() {

	var s Student3
	s.score = 90

	s.Person3 = new(Person3)
	s.Person3.name = "李寧"
	//(*s.Person3).name = "周杰倫"
	s.Person3.id = 110
	s.Person3.age = 20
	fmt.Println(s.name, s.id, s.age)

	s.Person3 = &Person3{101, "林俊杰", 35}
	fmt.Println(s.name, s.id, s.age)
}
