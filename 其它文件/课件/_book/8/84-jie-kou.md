### **接口**

在讲解具体的接口之前，先看如下问题。

使用面向对象的方式，设计一个加减的计算器

代码如下：

```go
package main

import "fmt"

//父类
type Operate struct {
	num1 int
	num2 int
}

//加法子类
type Add struct {
	Operate
}

//减法子类
type Sub struct {
	Operate
}

//加法子类的方法
func (a *Add) Result() int {
	return a.num1 + a.num2
}

//减法子类的方法
func (s *Sub) Result() int {
	return s.num1 - s.num2
}

//方法调用
func main0201() {
	//创建加法对象
	//var a Add
	//a.num1 = 10
	//a.num2 = 20
	//v := a.Result()
	//fmt.Println(v)

	//创建减法对象
	var s Sub
	s.num1 = 10
	s.num2 = 20
	v := s.Result()
	fmt.Println(v)
}
```

以上实现非常简单，但是有个问题，在main\( \)函数中，当我们想使用减法操作时，创建减法类的对象，调用其对应的减法的方法。但是，有一天，系统需求发生了变化，要求使用加法，不再使用减法，那么需要对main\( \)函数中的代码，做大量的修改。将原有的代码注释掉，创建加法的类对象，调用其对应的加法的方法。有没有一种方法，让main\( \)函数，只修改很少的代码就可以解决该问题呢？有，要用到接下来给大家讲解的接口的知识点。

### 1.**什么是接口**

接口就是一种规范与标准，在生活中经常见接口，例如:笔记本电脑的USB接口，可以将任何厂商生产的鼠标与键盘，与电脑进行链接。为什么呢？原因就是，USB接口将规范和标准制定好后，各个生产厂商可以按照该标准生产鼠标和键盘就可以了。

在程序开发中，接口只是规定了要做哪些事情，干什么。具体怎么做，接口是不管的。这和生活中接口的案例也很相似，例如：USB接口，只是规定了标准，但是不关心具体鼠标与键盘是怎样按照标准生产的.

在企业开发中，如果一个项目比较庞大，那么就需要一个能理清所有业务的架构师来定义一些主要的接口，这些接口告诉开发人员你需要实现那些功能。

### 3.**接口定义**

接口定义的语法如下：

```go
//先定义接口  一般以er结尾  根据接口实现功能
type Humaner interface {
	//方法  方法的声明
	sayhi()
}
```

怎样具体实现接口中定义的方法呢？

```go
type student11 struct {
	name  string
	age   int
	score int
}

func (s *student11)sayhi()  {
	fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

type teacher11 struct {
	name    string
	age     int
	subject string
}

func (t *teacher11)sayhi()  {
	fmt.Printf("大家好，我是%s,今年%d岁，我的学科是%s\n",t.name,t.age,t.subject)
}
```

具体的调用如下：

```go
func main() {
	//接口是一种数据类型 可以接收满足对象的信息
	//接口是虚的  方法是实的
	//接口定义规则  方法实现规则
	//接口定义的规则  在方法中必须有定义的实现
	var h Humaner

	stu := student11{"小明",18,98}
	//stu.sayhi()
	//将对象信息赋值给接口类型变量
	h = &stu
	h.sayhi()

	tea := teacher11{"老王",28,"物理"}
	//tea.sayhi()
	//将对象赋值给接口 必须满足接口中的方法的声明格式
	h = &tea
	h.sayhi()
}
```

只要类\(结构体\)实现对应的接口，那么根据该类创建的对象，可以赋值给对应的接口类型。

接口的命名习惯以er结尾。

### 3.**多态**

接口有什么好处呢？实现多态。

多态就是同一个接口，使用不同的实例而执行不同操作

所谓多态指的是多种表现形式，如下图所示：

![](/assets/import13.png)

使用接口实现多态的方式如下：

```go
package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner1 interface {
    //方法  方法的声明
    sayhi()

}


type student12 struct {
    name  string
    age   int
    score int
}

func (s *student12)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

type teacher12 struct {
    name    string
    age     int
    subject string
}

func (t *teacher12)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的学科是%s\n",t.name,t.age,t.subject)
}

//多态的实现
//将接口作为函数参数  实现多态
func sayhello(h Humaner1)  {
    h.sayhi()
}

func main() {

    stu := student12{"小明",18,98}
    //调用多态函数
    sayhello(&stu)

    tea := teacher12{"老王",28,"Go"}
    sayhello(&tea)
}
```

关于接口的定义，以及使用接口实现多态，大家都比较熟悉了，但是多态有什么好处呢？现在还是以开始提出的计算器案例给大家讲解一下，在开始我们已经实现了一个加减功能的计算器，但是有同学感觉太麻烦了，因为实现加法，就要定义加法操作的类（结构体）,实现减法就要定义减法的类\(结构体\)，所以该同学实现了一个比较简单的加减法的计算器，如下所示：

1.使用面向对象的思想实现一个加减功能的计算器，可能有同学感觉非常简单，代码如下：

![](file:///C:\Users\HP\AppData\Local\Temp\ksohtml\wps92B.tmp.jpg)

我们定义了一个类\(结构体\)，然后为该类创建了一个方法，封装了整个计算器功能，以后要使用直接使用该类\(结构体\)创建对象就可以了。这就是面向对象总的封装性。

也就是说，当你写完这个计算器后，交给你的同事，你的同事要用，直接创建对象，然后调用GetResult\(\)方法就可以，根本不需要关心该方法是怎样实现的.这不是我们前面在讲解面向对象概念时说到的，找个对象来干活吗？不需要自己去实现该功能。

2.大家仔细观察上面的代码，有什么问题吗？

现在让你在改计算器中，再增加一个功能，例如乘法，应该怎么办呢？你可能会说很简单啊，直接在GetResult\( \)方法的switch中添加一个case分支就可以了。

问题是：在这个过程中，如果你不小心将加法修改成了减法怎么办？或者说，对加法运算的规则做了修改怎么办？举例子说明：

你可以把该程序方法想象成公司中的薪资管理系统。如果公司决定对薪资的运算规则做修改，由于所有的运算规则都在Operation类中的GetResult\(\)方法中，所以公司只能将该类的代码全部给你，你才能进行修改。这时，你一看自己作为开发人员工资这么低，心想“TMD,老子累死累活才给这么点工资，这下有机会了”。直接在自己工资后面加了3000

numA+numB+3000

所以说，我们应该将加减等运算分开，不应该全部糅合在一起，这样你修改加的时候，不会影响其它的运算规则：

具体实现如下：

![](file:///C:\Users\HP\AppData\Local\Temp\ksohtml\wps92C.tmp.jpg)

现在已经将各个操作分开了，并且这里我们还定义了一个父类\(结构体\)，将公共的成员放在该父类中。如果现在要修改某项运算规则，只需将对应的类和方法发给你，进行修改就可以了。

这里的实现虽然将各个运算分开了，但是与我们第一次实现的还是有点区别。我们第一次实现的加减计算器也是将各个运算分开了，但是没有定义接口。那么该接口的意义是什么呢?继续看下面的问题。

3：现在怎样调用呢？

这就是我们一开始给大家提出的问题，如果调用的时候，直接创建加法操作的对象，调用对应的方法，那么后期要改成减法呢？需要做大量的修改，所以问题解决的方法如下：

![](file:///C:\Users\HP\AppData\Local\Temp\ksohtml\wps92D.tmp.jpg)

创建了一个类OperationFactory，在改类中添加了一个方法CreateOption\( \)负责创建对象，如果输入的是“+”，创建

OperationAdd的对象，然后调用OperationWho\( \)方法，将对象的地址传递到该方法中，所以变量i指的就是OperationAdd，接下来在调用GetResult\( \)方法，实际上调用的是OperationAdd类实现的GetResult\( \)方法。

同理如果传递过来的是“-”，流程也是一样的。

所以，通过该程序，大家能够体会出多态带来的好处。

4：最后调用

![](file:///C:\Users\HP\AppData\Local\Temp\ksohtml\wps92E.tmp.jpg)

这时会发现调用，非常简单，如果现在想计算加法，只要将”-”,修改成”+”就可以。也就是说,除去了main\( \)函数与具体运算类的依赖。

当然程序经过这样设计以后：如果现在修改加法的运算规则，只需要修改OperationAdd类中对应的方法，

不需要关心其它的类，如果现在要增加“乘法” 功能，应该怎样进行修改呢？第一：定义乘法的类，完成乘法运算。

第二：在OperationFactory类中CrateOption\( \)方法中添加相应的分支。但是这样做并不会影响到其它的任何运算。

大家可以自己尝试实现“乘法”与“除法”的运算。

在使用面向对象思想解决问题时，一定要先分析，定义哪些类，哪些接口，哪些方法。把这些分析定义出来，然后在考虑具体实现。

最后完整代码如下：

```go
package main

import "fmt"

//定义接口
type Opter interface {
    //方法声明
    Result() int
}

//父类
type Operate struct {
    num1 int
    num2 int
}
//加法子类
type Add struct {
    Operate
}

//加法子类的方法
func (a *Add) Result() int {
    return a.num1 + a.num2
}

//减法子类
type Sub struct {
    Operate
}

//减法子类的方法
func (s *Sub) Result() int {
    return s.num1 - s.num2
}

//创建一个类负责对象创建
//工厂类
type Factory struct {
}

func (f *Factory) Result(num1 int, num2 int, ch string) {
    switch ch {
    case "+":
        var a Add
        a.num1 = num1
        a.num2 = num2
        Result(&a)
    case "-":
        var s Sub
        s.num1 = num1
        s.num2 = num2
        Result(&s)

    }
}

//通过设计模式调用
func main() {
    //创建工厂对象
    var f Factory
    f.Result(10, 20, "+")
}
```

下面我们将接口其它的知识点再给大家说一下：

### 4.**接口继承与转换（了解）**

接口也可以实现继承:

```go
package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner2 interface {   //子集
    //方法  方法的声明
    sayhi()

}

type Personer interface {  //超集
    Humaner2   //继承sayhi()

    sing(string)
}

type student13 struct {
    name  string
    age   int
    score int
}

func (s *student13)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

func (s *student13)sing(name string)  {
    fmt.Println("我为大家唱首歌",name)
}

func main() {
    //接口类型变量定义
    var h Humaner2
    var stu student13 = student13{"小吴",18,59}
    h = &stu
    h.sayhi()

    //接口类型变量定义
    var p Personer
    p = &stu
    p.sayhi()
    p.sing("大碗面")
}
```

接口继承后，可以实现“超集”接口转换“子集”接口，代码如下：

```go
package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner2 interface {   //子集
    //方法  方法的声明
    sayhi()

}

type Personer interface {  //超集
    Humaner2   //继承sayhi()

    sing(string)
}

type student13 struct {
    name  string
    age   int
    score int
}

func (s *student13)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

func (s *student13)sing(name string)  {
    fmt.Println("我为大家唱首歌",name)
}

func main()  {
    //接口类型变量定义
    var h Humaner2  //子集
    var p Personer    //超集
    var stu student13 = student13{"小吴",18,59}

    p = &stu
    //将一个接口赋值给另一个接口
    //超集中包含所有子集的方法
    h = p  //ok

    h.sayhi()

    //子集不包含超集
    //不能将子集赋值给超集
    //p = h  //err
    //p.sayhi()
    //p.sing("大碗面")
}
```

### 5.**空接口**

空接口\(interface{}\)不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。

例如：

```go
var i interface{}
//接口类型可以接收任意类型的数据
//fmt.Println(i)
fmt.Printf("%T\n",i)
i = 10
fmt.Println(i)
fmt.Printf("%T\n",i)
```

当函数可以接受任意的对象实例时，我们会将其声明为interface{}，最典型的例子是标准库fmt中PrintXXX系列的函数，例如：

```
func Printf(fmt string, args ...interface{})
```

`func Println(args ...interface{})`

如果自己定义函数，可以如下：

```
func Test(arg ...interface{}) {

}
```

Test\( \)函数可以接收任意个数，任意类型的参数。

