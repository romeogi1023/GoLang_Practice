## **写入数据**

文件打开以后，可以向文件中写数据，可以使用WriteString\( \)方法。

```go
//\反斜杠 转义字符
//在写路径时可以使用/正斜杠代替\反斜杠
fp,err := os.Create("D:/a.txt")
if err!=nil{
    //文件创建失败
    /*
    1.路径不存在
    2.文件权限
    3.程序打开文件上限
     */
    fmt.Println("文件创建失败")
    return
}

//写文件

//\n不会换行  原因 在windows文本文件中换行\r\n  回车  在linux中换行\n
fp.WriteString("hello world\r\n")
fp.WriteString("性感荷官在线发牌")


defer fp.Close()

//关闭文件
//如果打开文件不关闭 造成内存的浪费  程序打开文件的上限
//fp.Close()
```

WriteString\( \)方法默认返回两个参数，

```go
count,err1 := fp.WriteString("性感老王在线授课")

if err1!=nil {
    fmt.Println("写入文件失败")
    return
}else {
    fmt.Println(count)
}
```

第一个参数，指的是写入文件的数据长度，第二个参数记录的是错误信息。

WriteString\( \)方法默认写到文件中的数据是不换行的。如果想换行，可以采用如下的方式

```go
//\n不会换行  原因 在windows文本文件中换行\r\n  回车  在linux中换行\n
fp.WriteString("hello world\r\n")
fp.WriteString("性感荷官在线发牌")
```

除了使用WriteString\( \)函数向文件中写入数据以外，还可以使用Write\( \)函数，如下所示：

```go
fp,err := os.Create("D:/a.txt")
if err!=nil{
    //文件创建失败
    /*
    1.路径不存在
    2.文件权限
    3.程序打开文件上限
     */
    fmt.Println("文件创建失败")
    return
}

//写操作
//slice := []byte{'h','e','l','l','o'}
//count,err1 := fp.Write(slice)
count,err1 := fp.Write([]byte("性感老王在线授课"))

if err1!=nil {
    fmt.Println("写入文件失败")
    return
}else {
    fmt.Println(count)
}

defer fp.Close()
```

在这里要注意的是，使用Write\( \)函数写数据时，参数为字节切片，所以需要将字符串转换成字节切片。该方法返回的也是写入文件数据的长度

第三种写入的方式使用WriteAt\( \)函数，在指定的位置写入数据

```go
fp,err := os.Create("D:/a.txt")
if err!=nil{
    //文件创建失败
    /*
    1.路径不存在
    2.文件权限
    3.程序打开文件上限
     */
    fmt.Println("文件创建失败")
    return
}

//写操作
//获取光标流位置'
//获取文件起始到结尾有多少个字符
//count,_:=fp.Seek(0,os.SEEK_END)
count,_:=fp.Seek(0,io.SeekEnd)
fmt.Println(count)
//指定位置写入
fp.WriteAt([]byte("hello world"),count)
fp.WriteAt([]byte("hahaha"),0)
fp.WriteAt([]byte("秀儿"),19)

defer fp.Close()
```

以上程序中Seek\( \)函数返回值存储到变量n中，值为文件末尾的位置。WriteAt\( \)也返回的是写入的数据长度。

以上就是我们常用的关于向文件中写入数据的方式，但是有同学可能有疑问，每次向文件中写入数据之前，都是先执行了，Create\( \)这个函数，而这个函数的作用前面我们也已经说过，有两个作用：第一：创建新文件。第二：如果所创建的文件已经存在，会删除掉文件中存储的数据。那么，现在怎样向已有的文件中追加数据呢？

如果要解决这个问题，那么大家一定要注意的就是，对已经存在的文件不能再执行Create\( \),而是要执行OpenFile\( \).

如下所示：

```go
//os.Open  只读方式打开
//fp,err := os.Open("D:/a.txt")

//os.OpenFile(文件名，打开方式，打开权限)
fp,err := os.OpenFile("D:/a.txt",os.O_RDWR,6)
if err!=nil {
    fmt.Println("打开文件失败")
}

fp.WriteString("hello")
fp.WriteAt([]byte("hello"),25)


defer fp.Close()
```

OpenFile\( \)这个函数有三个参数，第一个参数表示打开文件的路径，第二个参数表示模式，常见的模式有

O\_RDONLY\(只读模式\)，O\_WRONLY\(只写模式\),O\_RDWR\(可读可写模式\)，O\_APPEND\(追加模式\)。

第三个参数，表示权限，取值范围（0-7）

表示如下：

0：没有任何权限

1：执行权限\(如果是可执行文件，是可以运行的\)

2：写权限

3:写权限与执行权限

4：读权限

5:读权限与执行权限

6:读权限与写权限

7:读权限，写权限，执行权限

