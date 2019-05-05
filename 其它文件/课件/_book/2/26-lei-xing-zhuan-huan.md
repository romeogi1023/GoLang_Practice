## **2.6类型转换**

Go语言中不允许隐式转换，所有类型转换必须显式声明，而且转换只能发生在两种相互兼容的类型之间。

```go
    var ch byte = 97
    //var a int = ch //err, cannot use ch (type byte) as type int in assignment
    var a int = int(ch)
```



