## **recover**

运行时panic异常一旦被引发就会导致程序崩溃。这当然不是我们愿意看到的，因为谁也不能保证程序不会发生任何运行时错误。

Go语言为我们提供了专用于“拦截”运行时panic的内建函数——recover。它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。

* **注意：recover只有在defer调用的函数中有效。**

示例如下：

```go
package main

import "fmt"

func testA()  {
	fmt.Println("testA")

}

func testB(x int)  {
	//设置recover()

	//在defer调用的函数中使用recover()
	defer func() {
		//防止程序崩溃
		recover()
	}()  //匿名函数

	var a [3]int
	a[x] = 999
}

func testC()  {
	fmt.Println("testC")
}
func main() {
	testA()
	testB(3)  //发生异常 中断程序
	testC()
}
```

以上程序的运行结果如下：

```
testA
testC
```

通过以上程序，我们发现虽然TestB\( \)函数会导致整个应用程序崩溃，但是由于在改函数中调用了recover\( \)函数，所以整个函数并没有崩溃。虽然程序没有崩溃，但是我们也没有看到任何的提示信息，那么怎样才能够看到相应的提示信息呢？

可以直接打印recover\( \)函数的返回结果，如下所示：

```go
func testB(x int)  {
	//设置recover()

	//在defer调用的函数中使用recover()
	defer func() {
		//防止程序崩溃
		//recover()
		fmt.Println(recover())    //直接打印
	}()  //匿名函数

	var a [3]int
	a[x] = 999
}
```

输出结果如下：

```
testA
runtime error: index out of range
testC
```

从输出结果发现，确实打印出了相应的错误信息。

但是，如果程序没有出错，也就是数组下标没有越界，会出现什么情况呢？

```go
func testA()  {
	fmt.Println("testA")

}
func testB(x int)  {
	//设置recover()

	//在defer调用的函数中使用recover()
	defer func() {
		//防止程序崩溃
		//recover()
		fmt.Println(recover())
	}()  //匿名函数

	var a [3]int
	a[x] = 999
}

func testC()  {
	fmt.Println("testC")
}
func main() {
	testA()
	testB(0)  //发生异常 中断程序
	testC()
}
```

输入的结果如下：

```
testA
<nil>
testC
```

这时输出的是空，但是我们希望程序没有错误的时候，不输出任何内容。



所以，程序修改如下：

```go
func testA()  {
	fmt.Println("testA")

}

func testB(x int)  {
	//设置recover()

	//在defer调用的函数中使用recover()
	defer func() {
		//防止程序崩溃
		//recover()
		//fmt.Println(recover())

		if err:=recover();err!=nil {
			fmt.Println(err)
		}
	}()  //匿名函数

	var a [3]int
	a[x] = 999
}

func testC()  {
	fmt.Println("testC")
}
func main() {
	testA()
	testB(0)  //发生异常 中断程序
	testC()
}

```

通过以上代码，发现其实就是加了一层判断。



