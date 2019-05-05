### **2：创建同级目录**

（2.1）创建src目录，在改目录下创建go源码文件

（2.1.1）在项目文件夹下新建src目录，如下图所示：

![](/assets/import01.png)

在D盘Workspace目录下创建了src目录。

（2.1.2）在改目录下创建不同的go源码文件，如下图所示：

![](/assets/import02.png)

在src目录下创建main.go文件和test.go文件（注意：这个两个文件是在同一个目录下面，都是在src目录下面）

main.go 文件下的代码如下所示：

```go
package mian

import "fmt"

func main () {
    fmt.Println("main")
}
```

test.go 文件下的代码如下所示：

```
package mian   //必须与main.go必须是一个包

import "fmt"

func Test () {
    fmt.Println("Test")
}
```

现在已经完成两个文件代码的编写，接下来的问题是，我们怎样在main.go文件中的入口函数main\( \)中调用test.go文件中的Test\( \)函数呢？这就需要设置环境变量GOPATH属性。如果要完成不同文件中函数的调用，必须设置GOPATH，否则，即使文件处于同一工作目录（工作区）下，也是无法完成调用的。

（2.2）GOPATH设置

GOPATH设置的具体步骤如下：

![](/assets/import50.png)

![](/assets/import51.png)

配置完成后，可以测试一下是否配置成功

![](/assets/import52.png)

（2.2） 在main.go文件中完成对test.go文件中函数的调用

![](/assets/import05.png)

最后编译执行。

注意：同一个目录下不能定义不同的package。

