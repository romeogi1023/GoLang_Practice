

## **延迟调用defer**

### **defer基本使用**

函数定义完成后，只有调用函数才能够执行，并且一经调用立即执行。例如：

```go
fmt.Println("hello")
fmt.Println("老王")
```

先输出“hello”,然后再输出“老王”

但是关键字 defer⽤于延迟一个函数（或者当前所创建的匿名函数）的执行。注意，defer语句只能出现在函数的内部。

基本用法如下：

```go
defer fmt.Println("hello")
fmt.Println("老王")
```

以上两行代码，输出的结果为，先输出“老王”，然后输出”hello”.

defer的应用场景：文件操作，先打开文件，执行读写操作，最后关闭文件。为了保证文件的关闭能够正确执行，可以使用defer.

### **defer执行顺序**

先看如下程序执行结果是：

```go
defer fmt.Println("hello")
defer fmt.Println("老王")
defer fmt.Println("你好")
```

执行的结果是：

`你好`

`老王`

`hello`

总结：如果一个函数中有多个defer语句，它们会以后进先出的顺序执行。

如下程序执行的结果：

```
func test03(x int) {
	v := 100 / x
	fmt.Println(v)
}
defer fmt.Println("hello")
defer fmt.Println("老王")
defer test03(0)
defer fmt.Println("你好")
```

执行结果：

```go
你好
老王
hello
panic: runtime error: integer divide by zero
```

即使函数或某个延迟调用发生错误，这些调用依旧会被执⾏。

### **defer与匿名函数结合使用**

我们先看以下程序的执行结果：

```go
a := 10
b := 20
defer func() {
	fmt.Println("匿名函数a", a)
	fmt.Println("匿名函数b", b)
}()

a = 100
b = 200
fmt.Println("main函数a", a)
fmt.Println("main函数b", b)
```

执行的结果如下：

```
main函数a 100
main函数b 200
匿名函数a 100
匿名函数b 200
```

前面讲解过，defer会延迟函数的执行，虽然立即调用了匿名函数，但是该匿名函数不会执行，等整个main\( \)函数结束之前在去调用执行匿名函数，所以输出结果如上所示。

现在将程序做如下修改：

```go
a := 10
b := 20
defer func(a,b int) {		//添加参数
	fmt.Println("匿名函数a", a)
	fmt.Println("匿名函数b", b)
}(a,b)     //传参

a = 100
b = 200
fmt.Println("main函数a", a)
fmt.Println("main函数b", b)
```

该程序的执行结果如下：

```
main函数a 100
main函数b 200
匿名函数a 10
匿名函数b 20
```

从执行结果上分析，由于匿名函数前面加上了defer所以，匿名函数没有立即执行。但是问题是，程序从上开始执行当执行到匿名函数时，虽然没有立即调用执行匿名函数，但是已经完成了参数的传递。



