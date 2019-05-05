package main

import "fmt"

func main() {

	score := 720

	//if score>=680 {
	//	fmt.Println("可以去清華")
	//	if score >= 680{
	//		fmt.Println("學汽修")
	//	}else if score<710 {
	//		fmt.Println("學美容美發")
	//	}else if score>710 {
	//		fmt.Println("學挖掘機")
	//	}
	//}else {
	//	fmt.Println("不能去清華")
	//}

	if score >= 680 {
		fmt.Println("可以去清華")
		if score >= 680 {
			fmt.Println("學汽修")
			if score > 700 {
				fmt.Println("學美容美發")
				if score > 710 {
					fmt.Println("學挖掘機")
				}
			}
		}

	} else {
		fmt.Println("不能去清華")
	}

}
