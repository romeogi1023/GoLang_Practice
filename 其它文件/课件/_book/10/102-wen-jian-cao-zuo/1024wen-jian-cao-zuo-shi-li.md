## **文件操作案例**

文件拷贝，将已有的文件复制一份，同时重新命名。

基本的思路：

\(1\)让用户输入要拷贝的文件的名称\(源文件\)以及目的文件的名称

\(2\)创建目的文件

\(3\)打开源文件，并且读取该文件中的内容

\(4\)将从源文件中读取的内容写到目的文件中。

```go
var srcFileName string
var dstFileName string
fmt.Printf("请输入源文件名称:")
fmt.Scan(&srcFileName)
fmt.Println("请输入目的文件名称:")
fmt.Scan(&dstFileName)
if srcFileName == dstFileName {
    fmt.Println("源文件和目的文件名字不能相同")
    return
}
//只读方式打开源文件
sF,err1 := os.Open(srcFileName)
if err1 != nil {
    fmt.Println("err1=",err1)
    return
}
//新建目的文件
dF,err2 := os.Create(dstFileName)
if err2 != nil{
    fmt.Println("err2=",err2)
    return
}
//操作完毕，需要关闭文件
defer sF.Close()
defer dF.Close()
//核心处理，从源文件读取内容，往目的文件写，读多少写多少
buf := make([]byte,4*1024)//4k大小临时缓冲区
for{
    n,err := sF.Read(buf)//从源文件读取内容,每次读取一部分
    if err != nil{
        fmt.Println("err=",err)
        if err == io.EOF{//文件读取完毕
            break
        }
    }
    //往目的文件写，读多少写多少
    dF.Write(buf[:n])
}
```



