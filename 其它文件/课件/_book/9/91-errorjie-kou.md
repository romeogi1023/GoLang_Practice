## **error接口**

Go语言引入了一个关于错误处理的标准模式，即error接口，它是Go语言内建的接口类型，该接口的定义如下：

```go
type error interface {
    Error() string
}
```

Go语言的标准库代码包errors为用户提供如下方法：

![](/assets/import92.png)

通过以上代码，可以发现error接口的使用是非常简单的（error是一个接口，该接口只声明了一个方法Error\(\)，返回值是string类型，用以描述错误

）。下面看一下基本使用,

首先导包：

```go
import "errors"
```

然后调用其对应的方法：

![](/assets/import93.png)

当然fmt包中也封装了一个专门输出错误信息的方法，如下所示：

![](/assets/import94..png)

了解完基本的语法以后，接下来使用error接口解决Test\( \)函数被0整除的问题，如下所示：

![](/assets/impor95.png)

在Test\( \)函数中，判断变量b的取值，如果有误，返回错误信息。并且在main\( \)中接收返回的错误信息，并打印出来。

这种用法是非常常见的，例如，后面讲解到文件操作时，涉及到文件的打开，如下：

![](/assets/import96.png)

在打开文件时，如果文件不存在，或者文件在磁盘上存储的路径写错了，都会出现异常，这时可以使用error记录相应的错误信息。

