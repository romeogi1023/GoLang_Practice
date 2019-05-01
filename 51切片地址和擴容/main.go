package main

import (
	"fmt"
)

func main51() {
	//空切片 指向內存地址吧編號為0的空間
	var slice []int
	fmt.Printf("%p\n", slice)
	//當使用append進行追加數據時，切片地址可能會發生改變
	slice = append(slice, 1, 2, 3)
	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", &slice[0])
	fmt.Printf("%p\n", &slice[1])

	slice = append(slice, 4, 5, 6)
	fmt.Printf("%p\n", slice)
}

func main() {
	//容量
	s := make([]int, 5, 8)
	fmt.Println(s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	s1 := make([]int, 0, 1)
	oldcap := cap(s1)
	for i := 0; i < 200000; i++ {
		s1 = append(s1, i)
		newcap := cap(s1)
		if oldcap < cap(s1) {
			fmt.Printf("cap:%d  ===     %d\n", oldcap, newcap)
			oldcap = newcap
		}
	}
	//在使用append追加數據時，長度超過容量，容量自動擴容
	//一般 容量*2   如果超過 1024字節   每次擴容上一次的1/4
	//容量擴容每次偶數
}
