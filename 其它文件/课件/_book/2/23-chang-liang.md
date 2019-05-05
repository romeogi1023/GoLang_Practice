## **2.3常量**

在Go语言中，常量是指编译期间就已知且不可改变的值。常量可以是数值类型（包括整型、浮点型和复数类型）、布尔类型、字符串类型等。

### **2.3.1字面常量\(常量值\)**

所谓字面常量（literal），是指程序中硬编码的常量，如：

```
123
3.1415  // 浮点类型的常量
3.2+12i // 复数类型的常量
true  // 布尔类型的常量
"foo" // 字符串常量
```

### **2.3.2常量定义**

```go
const Pi float64 = 3.14
const zero = 0.0 // 浮点常量, 自动推导类型

const ( // 常量组
    size int64 = 1024
    eof        = -1 // 整型常量, 自动推导类型
)
const u, v float32 = 0, 3 // u = 0.0, v = 3.0，常量的多重赋值
const a, b, c = 3, 4, "foo"
// a = 3, b = 4, c = "foo"    //err, 常量不能修改
```

### **2.3.3iota枚举**

常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。

**在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。**

```go
const (
    x = iota // x == 0
    y = iota // y == 1
    z = iota // z == 2
    w  // 这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
    h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
    a       = iota //a=0
    b       = "B"
    c       = iota             //c=2
    d, e, f = iota, iota, iota //d=3,e=3,f=3
    g       = iota             //g = 4
)

const (
    x1 = iota * 10 // x1 == 0
    y1 = iota * 10 // y1 == 10
    z1 = iota * 10 // z1 == 20
)
```



