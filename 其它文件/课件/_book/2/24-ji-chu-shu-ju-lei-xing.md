

## **2.4基础数据类型**

### **2.4.1分类**

Go语言内置以下这些基础类型：

| **类型** | **名称** | **长度** | **零值** | **说明** |
| :--- | :--- | :--- | :--- | :--- |
| bool | 布尔类型 | 1 | false | 其值不为真即为家，不可以用数字代表true或false |
| byte | 字节型 | 1 | 0 | uint8别名 |
| rune | 字符类型 | 4 | 0 | 专用于存储unicode编码，等价于uint32 |
| int, uint | 整型 | 4或8 | 0 | 32位或64位 |
| int8, uint8 | 整型 | 1 | 0 | -128 ~ 127, 0 ~ 255 |
| int16, uint16 | 整型 | 2 | 0 | -32768 ~ 32767, 0 ~ 65535 |
| int32, uint32 | 整型 | 4 | 0 | -21亿~ 21亿, 0 ~ 42亿 |
| int64, uint64 | 整型 | 8 | 0 |  |
| float32 | 浮点型 | 4 | 0.0 | 小数位精确到7位 |
| float64 | 浮点型 | 8 | 0.0 | 小数位精确到15位 |
| complex64 | 复数类型 | 8 |  |  |
| complex128 | 复数类型 | 16 |  |  |
| uintptr | 整型 | 4或8 |  | ⾜以存储指针的uint32或uint64整数 |
| string | 字符串 |  | "" | utf-8字符串 |



### **2.4.2布尔类型**

```go
var v1 bool
v1 = true
v2 := (1 == 2) // v2也会被推导为bool类型

//布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
var b bool
b = 1 // err, 编译错误
b = bool(1) // err, 编译错误
```



### **2.4.3整型**

```go
varv1int32
v1=123
v2:=64//v1将会被自动推导为int类型
```

### 

### **2.4.4浮点型**

```go
 var f1 float32
    f1 = 12
    f2 := 12.0 // 如果不加小数点， fvalue2会被推导为整型而不是浮点型，float64
```

### 

### **2.4.5字符类型**

在Go语言中支持两个字符类型，一个是byte（实际上是uint8的别名），代表utf-8字符串的单个字节的值；另一个是rune，代表单个unicode字符。

```go
package main

import (
    "fmt"
)

func main() {
    var ch1, ch2, ch3 byte
    ch1 = 'a'  //字符赋值
    ch2 = 97   //字符的ascii码赋值
    ch3 = '\n' //转义字符
    fmt.Printf("ch1 = %c, ch2 = %c, %c", ch1, ch2, ch3)
}
```



### **2.4.6字符串**

在Go语言中，字符串也是一种基本类型：

```go
    var str string                                    // 声明一个字符串变量
    str = "abc"                                       // 字符串赋值
    ch := str[0]                                      // 取字符串的第一个字符
    fmt.Printf("str = %s, len = %d\n", str, len(str)) //内置的函数len()来取字符串的长度
    fmt.Printf("str[0] = %c, ch = %c\n", str[0], ch)

    //`(反引号)括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
    str2 := `hello
    mike \n \r测试
    `
    fmt.Println("str2 = ", str2)
    /*
        str2 =  hello
              mike \n \r测试
    */
```



### **2.4.7复数类型**

复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag）。

```go
    var v1 complex64 // 由2个float32构成的复数类型
    v1 = 3.2 + 12i
    v2 := 3.2 + 12i        // v2是complex128类型
    v3 := complex(3.2, 12) // v3结果同v2

    fmt.Println(v1, v2, v3)
    //内置函数real(v1)获得该复数的实部
    //通过imag(v1)获得该复数的虚部
    fmt.Println(real(v1), imag(v1))
```



