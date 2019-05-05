## **5.5 匿名函数与闭包**

所谓闭包就是一个函数“捕获”了和它在同一作用域的其它常量和变量。这就意味着当闭包被调用的时候，不管在程序什么地方调用，闭包能够使用这些常量或者变量。它不关心这些捕获了的变量和常量是否已经超出了作用域，所以只有闭包还在使用它，这些变量就还会存在。



在Go语言里，所有的匿名函数\(Go语言规范中称之为函数字面量\)都是闭包。匿名函数是指不需要定义函数名的一种函数实现方式，它并不是一个新概念，最早可以回溯到1958年的Lisp语言。



```go
func main() {
    i := 0
    str := "mike"

    //方式1
    f1 := func(){//匿名函数，无参无返回值
        //引用到函数外的变量
        fmt.Printf("方式1：i=%d,str=%s\n",i,str)
    }

    f1()//函数调用

    //方式1的另一种方式
    type FuncType func()    //声明函数类型,无参无返回值
    var f2 FuncType = f1
    f2()    //函数调用

    //方式2
    var f3 FuncType = func(){
        fmt.Printf("方式2：i=%d,str=%s\n",i,str)
    }
    f3()//函数调用

    //方式3
    func(){  //匿名函数，无参无返回值
        fmt.Printf("方式3：i=%d,str=%s\n",i,str)
}()    //别忘了后面的(),()的作用是，此处直接调用此匿名函数

    //方式4,匿名函数，有参有返回值
    v := func(a,b int) (result int) {
        result=a+b
        return
}(1,1)//别忘了后面的(1,1),(1,1)的作用是，此处直接调用此匿名函数，并传参
    fmt.Println("v=",v)

}
```



闭包捕获外部变量特点：

```go
func main() {
    i := 10
    str := "mike"
    func(){
        i = 100
        str = "go"
        //内部：i=100,str=go
        fmt.Printf("内部：i=%d,str=%s\n",i,str)
    }()    //别忘了后面的(),()的作用是，此处直接调用此匿名函数

    //外部：i=100,str=go
    fmt.Printf("外部：i=%d,str=%s\n",i,str)
}
```



函数返回值为匿名函数：

```go
//squares返回一个匿名函数，func()int
//该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
    var x int
    return func() int {//匿名函数
        x++//捕获外部变量
        return x*x
    }
}

func main() {
    f:=squares()
    fmt.Println(f())    //"1"
    fmt.Println(f())    //"4"
    fmt.Println(f())    //"9"
    fmt.Println(f())    //"16"
}

```

函数squares返回另一个类型为func\(\) int的函数。对squares的一次调用会生成一个局部变量x并返回一个匿名函数。每次调用时匿名函数时，该函数都会先使x的值加1，再返回x的平方。第二次调用squares时，会生成第二个x变量，并返回一个新的匿名函数。新匿名函数操作的是第二个x变量。



通过这个例子，我们看到变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于f中。



