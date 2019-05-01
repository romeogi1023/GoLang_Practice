package main

import "fmt"

func main() {

	//百錢百雞
	/*
		中國古代數學家張丘建在他的【算經】中提出了一個著名的"百錢百雞問題"：
		一只公雞值五錢，一只母雞值三錢，三只小雞值一錢，
		現在要用百錢買百雞，請問公雞、母雞、小雞各多少只？
	*/
	// x = 公雞個數、 y = 母雞個數、 z = 小雞個數

	//寫法一
	fmt.Println("寫法一")
	count1 := 0
	for x := 0; x <= 20; x++ {
		for y := 0; y <= 33; y++ {
			for z := 0; z <= 100; z += 3 {
				count1++
				if x+y+z == 100 && x*5+y*3+z/3 == 100 {
					fmt.Println(x, y, z)
				}
			}
		}
	}
	fmt.Println(count1)

	//寫法二
	fmt.Println("寫法二")
	count2 := 0
	for i := 0; i <= 20; i++ {
		for j := 0; j <= 33; j++ {
			count2++
			k := 100 - i - j
			if (k%3 == 0 && i*5+j*3+k/3 == 100) {
				fmt.Println(i, j, k)
			}
		}
	}
	fmt.Println(count2)
}
