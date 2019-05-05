## **5.3 递归函数**

递归指函数可以直接或间接的调用自身。



递归函数通常有相同的结构：一个跳出条件和一个递归体。所谓跳出条件就是根据传入的参数判断是否需要停止递归，而递归体则是函数自身所做的一些处理。



```go
//通过循环实现1+2+3……+100
func Test01() int {
    i := 1
    sum := 0
    for i = 1; i <= 100; i++ {
        sum += i
    }

    return sum
}

//通过递归实现1+2+3……+100
func Test02(num int) int {
    if num == 1 {
        return1
    }

    return num + Test02(num-1)//函数调用本身
}

//通过递归实现1+2+3……+100
func Test03(num int) int {
    if num == 100 {
        return 100
    }

    return num + Test03(num+1)  //函数调用本身
}

func main() {

    fmt.Println(Test01())    //5050
    fmt.Println(Test02(100))    //5050
    fmt.Println(Test03(1))    //5050
}

```



