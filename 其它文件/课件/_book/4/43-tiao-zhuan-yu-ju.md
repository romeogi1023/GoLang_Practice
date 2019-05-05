## **4.3跳转语句**

### **4.3.1break和continue**

在循环里面有两个关键操作break和continue，break操作是跳出当前循环，continue是跳过本次循环。

```go
for i := 0; i < 5; i++ {
    if 2 == i {
    //break//break操作是跳出当前循环
    continue//continue是跳过本次循环
    }
    fmt.Println(i)
}
```



注意：break可⽤于for、switch、select，⽽continue仅能⽤于for循环。



### **4.3.2goto**

用goto跳转到必须在当前函数内定义的标签：

```go
func main() {
    for i := 0; i < 5; i++{
        for{
            fmt.Println(i)
            gotoLABEL//跳转到标签LABEL，从标签处，执行代码
        }
    }

    fmt.Println("thisistest")

LABEL:
    fmt.Println("itisover")
}
```



