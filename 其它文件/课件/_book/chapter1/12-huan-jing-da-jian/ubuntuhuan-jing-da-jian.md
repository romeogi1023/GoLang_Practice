## GO-Linux系统安装并配置

### 第一步：

下载go的安装包，自动下载到下载目录

下载地址：

Go中文下载网：[https://www.golangtc.com/download](https://www.golangtc.com/download)

下载页面

![](file://C:/Users/HP/Desktop/%E7%9F%A5%E4%BA%86%E8%AF%BE%E5%A0%82/Snipaste_2018-12-21_14-09-07.png?lastModify=1548490566)

* 下载的文件在 Download 文件夹下面

1. 打开`Terminal`（终端）`Ctrl Alt T`

2. cd 到 刚刚下载的文件的路径下面 Ubuntu 默认是在 Download 文件夹下面

```
cd Download 
 //cd 下载 
```

1. 使用`tar`命令将 Go 安装包解压到`/usr/local`路径下

```
//这里面的 go1.10.linux-amd64.tar.gz 是下载的文件名，要换成你自己的


输入命令：


sudo tar -zxzf go1.10.linux-amd64.tar.gz  -C /usr/local
```

> 1.进入下载目录
>
> 2.输入sudo 表示使用管理员身份执行命令，需要输入密码

### 第二步：配置环境变量

一：需要先装vim

直接在终端执行以下命令：

```
sudo apt-get install vim
```

二：编辑$HOME/.profile文件

我们需要将/usr/local/go/bin添加到环境变量PATH中。可以通过vi直接将下面的内容添加到$HOME/.profile中（以.开头的文件是隐藏文件）

在`Terminal`中输入`vi $HOME/.profile`回车

![](file://C:/Users/HP/Desktop/%E7%9F%A5%E4%BA%86%E8%AF%BE%E5%A0%82/Snipaste_2018-12-21_14-57-45.png?lastModify=1548490566)

输入 “i” 出现插入， 就可以在文件末尾添加以下内容

```
export PATH=$PATH:/usr/local/go/bin 

export GOPATH=$HOME/go      //根据自己情况选择
```

配置GOPATH，就是存储go源文件的位置

​对于Ubuntu系统，默认使用Home/go目录作为GOPATH

​该目录下有三个子目录：src，pkg，bin

> GO代码必须在工作空间内。工作空间就是一个目录，其中包含三个子目录：
>
> ​src----里面每一个子目录，就是一个包，包内是go的源码文件
>
> ​pkg----编译后生成的，包的目标文件
>
> ​bin----生成的可执行的文件

保存配置：

按键盘上ESC键，然后输入 ” ：“ 冒号 ，再输入wq

> wq是保存并退出
>
> q 退出不保存

三：让配置文件生效

先进入usr/local目录下，打开终端，使用source命令让配置文件生效

```
source $HOME/.profile
```

在任意目录输入`go version`出现以下结果表示成功：

```
go version go1.9.2 linux/amd64
```

输出`go env`



![](file://C:/Users/HP/Desktop/%E7%9F%A5%E4%BA%86%E8%AF%BE%E5%A0%82/Snipaste_2018-12-21_15-17-00.png?lastModify=1548490566)

---

## 接着安装 Goland 开发工具（GoLand现在收费）

1. 官网下载地址：[https://www.jetbrains.com/go/](#)

![](/assets/import300.png)

下载页面

1. cd 到 刚刚下载的文件的路径下面 Ubuntu 默认是在 Download 文件夹下面

```
cd Download
```

1. 使用`tar`命令将 GoLand 安装包解压到`/opt`路径下

```
//这里面的 goland-2018.1.1.tar.gz 是下载的文件名，要换成你自己的
tar -zxzf goland-2018.1.1.tar.gz  -C /opt
```

1. 给脚本\(goland.sh\)执行权限

```
//进入脚本的所在目录。其中 Goland-2018.1.1 是我的解压文件夹名，要换成自己的。

cd /opt/Goland-2018.1.1/

sudo chmod a=+rx bin/idea.sh

//启动GoLand

sudo bin/idea.sh
```

1. 如果一切正常的话，就会弹出GoLand启动页面

   ![](/assets/import200.png)













