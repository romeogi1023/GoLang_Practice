## **8.2匿名组合**

### **8.2.1匿名字段**

一般情况下，定义结构体的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

当匿名字段也是一个结构体的时候，那么这个结构体所拥有的全部字段都被隐式地引入了当前定义的这个结构体。

```go
//人
type Person struct {
    name string
    sex byte
    age int
}
//学生
type Student struct {
    Person    //匿名字段，那么默认Student就包含了Person的所有字段
    id int
    addr string
}
```

### **8.2.2初始化**

```go
//人
type Person struct {
    name string
    sex byte
    age int
}
//学生
type Student struct {
    Person//匿名字段，那么默认Student就包含了Person的所有字段
    id int
    addr string
}
func main() {
    //顺序初始化
    s1 := Student{Person{"mike",'m',18},1,"sz"}
    //s1 = {Person:{name:mike sex:109 age:18}id:1 addr:sz}
    fmt.Printf("s1=%+v\n",s1)
    //s2 := Student{"mike",'m',18,1,"sz"}//err
    //部分成员初始化1
    s3 := Student{Person:Person{"lily",'f',19},id:2}
    //s3 = {Person:{name:lily sex:102 age:19}id:2 addr:}
    fmt.Printf("s3=%+v\n",s3)
    //部分成员初始化2
    s4 := Student{Person:Person{name:"tom"},id:3}
    //s4 = {Person:{name:tomsex:0age:0}id:3addr:}
    fmt.Printf("s4=%+v\n",s4)
}
```

### **8.2.3成员的操作**

```go
var s1 Student//变量声明
//给成员赋值
s1.name = "mike"//等价于s1.Person.name="mike"
s1.sex = 'm'
s1.age = 18
s1.id = 1
s1.addr = "sz"
fmt.Println(s1)    //{{mike 109 18}1 sz}
var s2 Student//变量声明
s2.Person = Person{"lily",'f',19}
s2.id = 2
s2.addr = "bj"
fmt.Println(s2)    //{{lily 102 19}2 bj}
```

### **8.2.4同名字段**

```go
//人
type Person struct{
    name string
    sex byte
    age int
}
//学生
type Student struct{
    Person    //匿名字段，那么默认Student就包含了Person的所有字段
    id int
    addr string
    name string    //和Person中的name同名
}
func main(){
    var s Student//变量声明
    //给Student的name，还是给Person赋值？
    s.name = "mike"
    //{Person:{name:sex:0age:0}id:0addr:name:mike}
    fmt.Printf("%+v\n",s)
    //默认只会给最外层的成员赋值
    //给匿名同名成员赋值，需要显示调用
    s.Person.name = "yoyo"
    //Person:{name:yoyosex:0age:0}id:0addr:name:mike}
    fmt.Printf("%+v\n",s)
}
```

### **8.2.5其它匿名字段**

8.2.5.1非结构体类型

所有的内置类型和自定义类型都是可以作为匿名字段的：

```go
type mystr string//自定义类型
type Person struct {
    name string
    sex byte
    age int
}
type Student struct {
    Person    //匿名字段，结构体类型
    int    //匿名字段，内置类型
    mystr    //匿名字段，自定义类型
}
func main() {
    //初始化
    s1 := Student{Person{"mike",'m',18},1,"bj"}
    //{Person:{name:mikesex:109age:18}int:1mystr:bj}
    fmt.Printf("%+v\n",s1)
    //成员的操作，打印结果：mike,m,18,1,bj
    fmt.Printf("%s,%c,%d,%d,%s\n",s1.name,s1.sex,s1.age,s1.int,s1.mystr)
}
```

8.2.5.2结构体指针类型

```go
type Person struct {    //人
    name string
    sex byte
    age int
}
type Student struct {//学生
    *Person    //匿名字段，结构体指针类型
    id int
    addr string
}
func main() {
    //初始化
    s1 := Student{&Person{"mike",'m',18},1,"bj"}
    //{Person:0xc0420023e0id:1addr:bj}
    fmt.Printf("%+v\n",s1)
    //mike,m,18
    fmt.Printf("%s,%c,%d\n",s1.name,s1.sex,s1.age)
    //声明变量
    var s2 Student
    s2.Person = new(Person)//分配空间
    s2.name = "yoyo"
    s2.sex = 'f'
    s2.age = 20
    s2.id = 2
    s2.addr = "sz"
    //yoyo10220220
    fmt.Println(s2.name,s2.sex,s2.age,s2.id,s2.age)
}
```



