package main

import (
	"fmt"
)

func demo57(m map[int]string) {
	m[4] = "黃月英"
	fmt.Printf("%p\n", m)
	fmt.Println(m)
}

func main() {

	m := map[int]string{1: "貂蟬", 2: "大喬", 3: "小喬"}
	//地址傳遞  引用傳遞  形參和實參指向相同的內存地址  修改形參 會影響實參的值
	demo57(m)
	fmt.Printf("%p\n", m)
	fmt.Println(m)
}
