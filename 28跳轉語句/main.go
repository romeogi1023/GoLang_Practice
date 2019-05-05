package main

import "fmt"

func main() {

	//i:=0
	//for {  //for後面不寫任何的東西  那麽就是壹個死循環  永遠為真
	//	i++
	//	time.Sleep(time.Second)
	//
	//	if i == 5{
	//		//break
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	//1.break  跳出循環   如果多層for循環  跳出最近的內循環

	//2.continue  跳出本次循環  進入下次循環
	//註意：continue只能在for循環使用

	//3.goto 跳轉  可以用在任何地方，但是不能跨函數使用

	fmt.Println("aaaaaaaa")
	//goto關鍵字  End 是標簽  用戶起的名字
	fmt.Println("bbbbbbbb")
End:
	fmt.Println("cccccccc")
	goto End

}
