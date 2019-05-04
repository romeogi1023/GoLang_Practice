package main

import "fmt"

func main8301() {

	var i interface{}
	i = 10.34

	//value, ok := element.(T)
	//值, 值的判斷 := 接口變量.(數據類型)

	value, ok := i.(int)
	if ok {
		fmt.Println("整型數據", value)
	} else {
		fmt.Println("不是整型", value)
	}

}

func demo83() {
	fmt.Println("hello")
}

//類型斷言
func main8302() {

	var i []interface{}
	i = append(i, 10, 3.14, "aaa", demo83)

	//fmt.Println(i)

	for _, v := range i {
		fmt.Println(v)

		if data, ok := v.(int); ok {
			fmt.Println("整型數據", data)
		} else if data, ok := v.(float64); ok {
			fmt.Println("浮點型數據", data)
		} else if data, ok := v.(string); ok {
			fmt.Println("字符串", data)
		} else if data, ok := v.(func()); ok {
			//函數調用
			data()
		}
	}
}

func main() {
	var i []interface{}
	i = append(i, 10, 3.14, "aaa", demo83)

	//fmt.Println(i)

	for _, v := range i {

		switch value := v.(type) {
		case int:
			fmt.Println("整型數據", value)
		case float64:
			fmt.Println("浮點型數據", value)
		case string:
			fmt.Println("字符串", value)
		case func():
			value()
		}

	}
}
