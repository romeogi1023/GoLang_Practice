## **8.3方法**

### **8.3.1概述**

在面向对象编程中，一个对象其实也就是一个简单的值或者一个变量，在这个对象中会包含一些函数，**这种带有接收者的函数，我们称为方法\(method\)**。本质上，一个方法则是一个和特殊类型关联的函数。

一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象，而是借助方法来做这些事情。

在Go语言中，可以给任意自定义类型（包括内置类型，但不包括指针类型）添加相应的方法。

⽅法总是绑定对象实例，并隐式将实例作为第⼀实参 \(receiver\)，方法的语法如下：

```go
func (receiver ReceiverType) funcName (parameters) (results)
```

* 参数 receiver 可任意命名。如⽅法中未曾使⽤，可省略参数名。
* 参数 receiver 类型可以是 T 或 \*T。基类型 T 不能是接⼝或指针。
* 不支持重载方法，也就是说，不能定义名字相同但是不同参数的方法。

### **8.3.2为类型添加方法**

8.3.2.1基础类型作为接收者

```go
type MyInt int//自定义类型，给int改名为MyInt

//在函数定义时，在其名字之前放上一个变量，即是一个方法
func (a MyInt) Add(b MyInt) MyInt {//面向对象
    return a + b
}

//传统方式的定义
func Add(a, b MyInt) MyInt {//面向过程
    return a + b
}

func main() {
    var a MyInt=1
    var b MyInt=1

//调用func (aMyInt) Add(bMyInt)
fmt.Println("a.Add(b)=",a.Add(b))//a.Add(b)=2

//调用func Add(a,bMyInt)
fmt.Println("Add(a,b)=",Add(a,b))//Add(a,b)=2
}
```

通过上面的例子可以看出，面向对象只是换了一种语法形式来表达。方法是函数的语法糖，因为receiver其实就是方法所接收的第1个参数。

注意：虽然方法的名字一模一样，但是如果接收者不一样，那么方法就不一样。

8.3.2.2结构体作为接收者

方法里面可以访问接收者的字段，调用方法通过点\(**.** \)访问，就像struct里面访问字段一样：

```go
type Person struct {
    name string
    sex byte
    age int
}

func (p Person) PrintInfo(){//给Person添加方法
    fmt.Println(p.name,p.sex,p.age)
}

func main() {
    p:=Person{"mike",'m',18}//初始化
    p.PrintInfo()//调用func(pPerson)PrintInfo()
}
```

### **8.3.3值语义和引用语义**

```go
type Person struct {
    name string
    sex byte
    age int
}

//指针作为接收者，引用语义
func (p *Person) SetInfoPointer(){
    //给成员赋值
    (*p).name = "yoyo"
    p.sex = 'f'
    p.age = 22
}

//值作为接收者，值语义
func (p Person) SetInfoValue(){
//给成员赋值
    p.name = "yoyo"
    p.sex = 'f'
    p.age = 22
}

func main() {
    //指针作为接收者，引用语义
    p1 := Person{"mike",'m',18}//初始化
    fmt.Println("函数调用前=",p1)//函数调用前={mike10918}
    (&p1).SetInfoPointer()
    fmt.Println("函数调用后=",p1)//函数调用后={yoyo10222}

    fmt.Println("==========================")

    p2 := Person{"mike",'m',18}//初始化
    //值作为接收者，值语义
    fmt.Println("函数调用前=",p2)//函数调用前={mike10918}
    p2.SetInfoValue()
    fmt.Println("函数调用后=",p2)//函数调用后={mike10918}
}
```

### **8.3.4方法集**

类型的方法集是指可以被该类型的值调用的所有方法的集合。

用实例实例 value 和 pointer 调用方法（含匿名字段）不受⽅法集约束，编译器编总是查找全部方法，并自动转换 receiver 实参。

8.3.4.1类型 \*T 方法集

一个指向自定义类型的值的指针，它的方法集由该类型定义的所有方法组成，无论这些方法接受的是一个值还是一个指针。

如果在指针上调用一个接受值的方法，Go语言会聪明地将该指针解引用，并将指针所指的底层值作为方法的接收者。

类型 \*T ⽅法集包含全部 receiver T + \*T ⽅法：

```go
type Person struct{
    name string
    sex byte
    age int
}

//指针作为接收者，引用语义
func (p *Person) SetInfoPointer(){
    (*p).name="yoyo"
    p.sex='f'
    p.age=22
}

//值作为接收者，值语义
func (p Person) SetInfoValue(){
    p.name="xxx"
    p.sex='m'
    p.age=33
}

func main() {
    //p为指针类型
    var p*Person = &Person{"mike",'m',18}
    p.SetInfoPointer()    //func (p)SetInfoPointer()

    p.SetInfoValue()    //func (*p)SetInfoValue()
    (*p).SetInfoValue()    //func (*p)SetInfoValue()
}
```

8.3.4.2类型 T 方法集

一个自定义类型值的方法集则由为该类型定义的接收者类型为值类型的方法组成，但是不包含那些接收者类型为指针的方法。

但这种限制通常并不像这里所说的那样，因为如果我们只有一个值，仍然可以调用一个接收者为指针类型的方法，这可以借助于Go语言传值的地址能力实现。

```go
type Person struct{
    name string
    sex byte
    age int
}

//指针作为接收者，引用语义
func (p *Person) SetInfoPointer(){
    (*p).name="yoyo"
    p.sex='f'
    p.age=22
}

//值作为接收者，值语义
func (p Person)SetInfoValue(){
    p.name="xxx"
    p.sex='m'
    p.age=33
}

func main() {
    //p为普通值类型
    var p Person = Person{"mike",'m',18}
    (&p).SetInfoPointer()    //func(&p)SetInfoPointer()
    p.SetInfoPointer()    //func(&p)SetInfoPointer()

    p.SetInfoValue()    //func(p)SetInfoValue()
    (&p).SetInfoValue()    //func(*&p)SetInfoValue()
}
```

### **8.3.5匿名字段**

8.3.5.1方法的继承

如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也能调用该方法。

```go
type Person struct {
    name string
    sex byte
    age int
}

//Person定义了方法
func (p *Person) PrintInfo() {
    fmt.Printf("%s,%c,%d\n",p.name,p.sex,p.age)
}

type Student struct {
    Person//匿名字段，那么Student包含了Person的所有字段
    id int
    addr string
}

func main() {
    p := Person{"mike",'m',18}
    p.PrintInfo()

    s := Student{Person{"yoyo",'f',20},2,"sz"}
    s.PrintInfo()
}
```

8.3.5.2方法的重写

```go
type Person struct {
    name string
    sex byte
    age int
}
//Person定义了方法
func (p *Person) PrintInfo() {
    fmt.Printf("Person:%s,%c,%d\n",p.name,p.sex,p.age)
}
type Student struct {
    Person//匿名字段，那么Student包含了Person的所有字段
    id int
    addr string
}
//Student定义了方法
func (s *Student) PrintInfo() {
    fmt.Printf("Student：%s,%c,%d\n",s.name,s.sex,s.age)
}
func main() {
    p:=Person{"mike",'m',18}
    p.PrintInfo()    //Person:mike,m,18
    s:=Student{Person{"yoyo",'f',20},2,"sz"}
    s.PrintInfo()    //Student：yoyo,f,20
    s.Person.PrintInfo()    //Person:yoyo,f,20
}
```

### **8.3.6方法值和方法表达式**

类似于我们可以对函数进行赋值和传递一样，方法也可以进行赋值和传递。

根据调用者不同，方法分为两种表现形式：方法值和方法表达式。两者都可像普通函数那样赋值和传参，区别在于方法值绑定实例，⽽方法表达式则须显式传参。

8.3.6.1方法值

```go
type Person struct{
    name string
    sex byte
    age int
}
func (p *Person) PrintInfoPointer() {
    fmt.Printf("%p,%v\n",p,p)
}
func (p Person) PrintInfoValue(){
    fmt.Printf("%p,%v\n",&p,p)
}
func main() {
    p:=Person{"mike",'m',18}
    p.PrintInfoPointer()    //0xc0420023e0,&{mike 109 18}
    pFunc1:=p.PrintInfoPointer    //方法值，隐式传递 receiver
    pFunc1()    //0xc0420023e0,&{mike 109 18}
    pFunc2:=p.PrintInfoValue
    pFunc2()    //0xc042048420,{mike 109 18}
}
```

8.3.6.2方法表达式

```go
type Person struct {
    name string
    sex byte
    age int
}
func (p *Person) PrintInfoPointer() {
    fmt.Printf("%p,%v\n",p,p)
}
func (p Person) PrintInfoValue() {
    fmt.Printf("%p,%v\n",&p,p)
}
func main() {
    p:=Person{"mike",'m',18}
    p.PrintInfoPointer()//0xc0420023e0,&{mike 109 18}
    //方法表达式，须显式传参
    //func pFunc1 (p *Person))
    pFunc1:=(*Person).PrintInfoPointer
    pFunc1(&p)    //0xc0420023e0,&{mike 109 18}
    pFunc2:=Person.PrintInfoValue
    pFunc2(p)    //0xc042002460,{mike 109 18}
}
```



