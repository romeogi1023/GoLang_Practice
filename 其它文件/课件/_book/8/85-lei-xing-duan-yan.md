## 类型查询

我们知道interface的变量里面可以存储任意类型的数值\(该类型实现了interface\)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：

* comma-ok断言
* switch测试

### 1.comma-ok断言

Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.\(T\)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。

如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。

```go
var i []interface{}

i = append(i, 10, 3.14, "aaa", demo15)

for _, v := range i {

    if data, ok := v.(int); ok {
        fmt.Println("整型数据：", data)
    } else if data, ok := v.(float64); ok {
        fmt.Println("浮点型数据：", data)
    } else if data, ok := v.(string); ok {
        fmt.Println("字符串数据：", data)
    } else if data, ok := v.(func()); ok {
        //函数调用
        data()
    }
}
```

### 2. switch测试

```go
var i []interface{}

i = append(i, 10, 3.14, "aaa", demo15)

for _,data := range i{
    switch value:=data.(type) {
    case int:
        fmt.Println("整型",value)
    case float64:
        fmt.Println("浮点型",value)
    case string:
        fmt.Println("字符串",value)
    case func():
        fmt.Println("函数",value)
    }
}
```



