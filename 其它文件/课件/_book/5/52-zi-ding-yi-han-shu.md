## **5.2 自定义函数**

### **5.2.1无参无返回值**

```go
func Test() {//无参无返回值函数定义
    fmt.Println("this is a test func")
}

func main() {
    Test() //无参无返回值函数调用
}
```

### **5.2.2有参无返回值**

5.2.2.1普通参数列表

```go
func Test01(v1int,v2int) {//方式1
    fmt.Printf("v1=%d,v2=%d\n",v1,v2)
}
func Test02(v1,v2int) {//方式2, v1, v2都是int类型
    fmt.Printf("v1=%d,v2=%d\n",v1,v2)
}
func main() {
    Test01(10,20) //函数调用
    Test02(11,22) //函数调用
}
```

5.2.2.2不定参数列表

1\)不定参数类型

不定参数是指函数传入的参数个数为不定数量。为了做到这点，首先需要将函数定义为接受不定参数类型：

```go
//形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数
func Test(args...int) {
    for_,n := range args{//遍历参数列表
        fmt.Println(n)
    }
}
func main() {
    //函数调用，可传0到多个参数
    Test()
    Test(1)
    Test(1,2,3)
}
```

2\)不定参数的传递

```go
func MyFunc01(args...int) {
    fmt.Println("MyFunc01")
    for_,n := range args{//遍历参数列表
        fmt.Println(n)
    }
}
func MyFunc02(args...int) {
    fmt.Println("MyFunc02")
    for_,n:=range args{//遍历参数列表
        fmt.Println(n)
    }
}
func Test(args...int) {
    MyFunc01(args...)//按原样传递,Test()的参数原封不动传递给MyFunc01
    MyFunc02(args[1:]...)//Test()参数列表中，第1个参数及以后的参数传递给MyFunc02
}
func main() {
    Test(1,2,3)//函数调用
}
```

### **5.2.3无参有返回值**

有返回值的函数，必须有明确的终止语句，否则会引发编译错误。

5.2.3.1一个返回值

```go
func Test01() int {//方式1
    return 250
}
//官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差
func Test02() (value int) {//方式2,给返回值命名
    value=250
    returnvalue
}
func Test03() (value int) {//方式3,给返回值命名
    value = 250
    return
}
funcmain(){
    v1 := Test01()//函数调用
    v2 := Test02()//函数调用
    v3 := Test03()//函数调用
    fmt.Printf("v1=%d,v2=%d,v3=%d\n", v1, v2, v3)
}
```

5.2.3.2多个返回值

```go
func Test01() (int,string) {//方式1
    return 250,"sb"
}
func Test02() (aint,str string) {//方式2,给返回值命名
    a=250
    str="sb"
    return
}
funcmain(){
    v1,v2:=Test01()//函数调用
    _,v3:=Test02()//函数调用，第一个返回值丢弃
    v4,_:=Test02()//函数调用，第二个返回值丢弃
    fmt.Printf("v1=%d,v2=%s,v3=%s,v4=%d\n",v1,v2,v3,v4)
}
```

### **5.2.4有参有返回值**

```go
//求2个数的最小值和最大值
func MinAndMax(num1 int, num2 int) (min int, max int) {
    if num1 > num2{//如果num1大于num2
        min = num2
        max = num1
    } else {
        max = num2
        min = num1
    }
    return
}
func main() {
    min,max := MinAndMax(33,22)
    fmt.Printf("min=%d,max=%d\n", min, max)  //min=22,max=33
}
```



