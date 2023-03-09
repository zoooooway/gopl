# 《Go程序设计语言》学习笔记

## 程序结构

### 包和文件
一个Go语言编写的程序对应一个或多个以.go为文件后缀名的源文件。
每个源文件中以包的声明语句开始，说明该源文件是属于哪个包。
包声明语句之后是 `import` 语句导入依赖的其它包。
然后是包一级的类型、变量、常量、函数的声明语句，包一级的各种类型的声明语句的顺序无关紧要。
> 函数内部的名字则必须先声明之后才能使用

#### import
`import` 语句支持导入其他包中导出的变量、常量、结构体、函数以及方法。
> 导出的成员的名称都是大写开头的，这是 `go` 语言的规则
>
注意：导入依赖 `go` 文件声明的 `package`，并且，你需要保证 `go` 文件中声明的 `package` 与所在文件夹名称一致。其次，需要导入本地包时，最好使用 *go.mod* 来进行管理。可以使用 `go mod init <module_name>` 来初始化 *go.mod* 文件。
```
# gopl
# │  ├── go.mod
# │  ├── handlers
# │  │  └── hello.go
# │  │  └── world.go
# │  └── main.go
```
如上文件结构，如果想要在 *main.go* 中导入 *handlers* 文件夹下 *hello.go* 中导出成员，那么 *hello.go* 中 `package` 的值必须是 handlers。



### 变量
简短变量声明语句中必须至少要声明一个新的变量.
```go
f, err := os.Open(infile)
// ...
f, err := os.Create(outfile) // compile error: no new variables
```
**简短变量声明语句只有对已经在同级词法域声明过的变量才和赋值操作语句等价，如果变量是在外部词法域声明的，那么简短变量声明语句将会在当前词法域重新声明一个新的变量**。
```go
in, err := os.Open(infile)
// ...
out, err := os.Create(outfile) // 此时，这条语句会给上面声明的err变量赋值

if true {
    err := "new err" // 此时这条语句会在当前词法域重新声明一个变量并赋值
}
```

#### 指针
如果用`var x int`声明语句声明一个x变量，那么`&x`表达式（取x变量的内存地址）将产生一个指向该整数变量的指针，指针对应的数据类型是`*int`，指针被称之为"指向int类型的指针"。如果指针名字为p，那么可以说"p指针指向变量x"，或者说"p指针保存了x变量的内存地址"。同时`*p`表达式对应p指针指向的变量的值。一般`*p`表达式读取指针指向的变量的值，这里为int类型的值，同时因为`*p`对应一个变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向的变量的值。
```go
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```
变量有时候被称为可寻址的值。即使变量由表达式临时生成，那么表达式也必须能接受&取地址操作。
```go
arr := [...]int{1, 2, 3, 4, 5}
s := arr[:]
reverseNew(&(arr[:])) // compile error: invalid operation: cannot take address of arr[:] (value of type []int)
reverseNew(&s)
```

#### 变量、指针与nil
Go 中的每个指针都有两个基本信息: 指针的类型和它所指向的值。我们将把它们表示为一个类似于(type, value)的一对值。
每个指针变量都需要一个类型，这就是为什么我们不能在不声明类型的情况下为变量赋一个 `nil` 值。
```go
// This does not work because we do not know the type
n := nil // compile error
```
观察下列代码:
```go
var a *int = nil
var b interface{} = nil

fmt.Printf("a=(%T, %v)\n", a, a) // a=(*int, <nil>)
fmt.Printf("b=(%T, %v)\n", b, b) // b=(<nil>, <nil>)
```
第二行输出似乎有点奇怪，似乎指针的类型应该是 `interface{}` 而不是 `nil`。
简而言之，就是因为我们使用了空接口，所以任何类型都会满足。`<nil>` 类型在技术上是一种特殊类型，它满足空接口，所以**当编译器不能确定其他类型信息时，就会使用`<nil>` 类型**。

观察下列代码:
```go
var a *int = nil
var b interface{} = a

fmt.Printf("a=(%T, %v)\n", a, a) // a=(*int, <nil>)
fmt.Printf("b=(%T, %v)\n", b, b) // b=(*int, <nil>)
```
可以看到，如果将a指针赋值给b，那么b指针就有了确定的类型信息，因此编译器不会再使用 `<nil>` 类型给b赋值。

理解上述所说后，再来看如下代码：
```go
var a *int = nil
var b interface{} = nil

// We will print out both type and value here
fmt.Printf("a=(%T, %v)\n", a, a) // a=(*int, <nil>)
fmt.Printf("b=(%T, %v)\n", b, b) // b=(<nil>, <nil>)
fmt.Println()
fmt.Println("a == nil:", a == nil) // true
fmt.Println("b == nil:", b == nil) // true
fmt.Println("a == b:", a == b) // false
```
你可能会觉得有些疑惑，为什么 `a == nil` 并且 `b == nil` ，但 `a != b`？
先理解 `==` 运算所比较的是什么会有助于理解这个现象。上述 `==` 运算实际可以表示成下列伪代码：
```go
a == nil: (*int, <nil>) == (*int*, <nil>)
b == nil: (<nil>, <nil>) == (<nil>, <nil>)
# Notice that these two are clearly not equal
# once we add in the type information.
a == b: (*int, <nil>) == (<nil>, <nil>)
```
也就是说，实际 `==` 运算时，不仅需要比较值，同时需要比较类型。
而为什么一个字面量形式的 `nil` 却能和不同类型的 `nil` 做比较呢？比如 `a == nil` 和 `b == nil`，a和b这两个变量的值虽然都是 `nil` ，但其类型却不同，在和字面量形式的 `nil` 做 `==` 运算时却都能得出相等的结果。
原因在于编译器，编译器会将 `nil` 强制转换为正确的类型，这类似于直接声明字面量整数时编译器将其转换为声明变量的类型：
```go
var a int = 12
var b float64 = 12
```
现在我们知道**在直接和字面量 `nil` 值作比较时，编译器会根据变量类型做强制转换**。

再来看看另一种情况:
```go
var a *int = nil
var b interface{} = a

// We will print out both type and value here
fmt.Printf("a=(%T, %v)\n", a, a) // a=(*int, <nil>)
fmt.Printf("b=(%T, %v)\n", b, b) // b=(*int, <nil>)
fmt.Println()
fmt.Println("a == nil:", a == nil) // true
fmt.Println("b == nil:", b == nil) // false
fmt.Println("a == b:", a == b) // true
```
也许你会再次感到困惑，为什么 `b == nil` 的结果是 `false` ?
上面我们提到：在直接和字面量 `nil` 值作比较时，编译器会根据变量类型做强制转换。那么在这里也许我们会想：b的实际类型是a的类型即 `*int`，那么在进行 `==` 运算时编译器应该将 `nil` 转换为 `(<nil>, <nil>)` 。
但实际上，这很难做到。在实际程序中，b变量的值可能发生多次改变——比如再次声明一个 `*string` 类型的变量并将其赋值给b。这意味着编译器没法在编译期确定b的实际类型，而只可能在运行期确定b的实际类型。这将有它自己的一套独特的复杂性，可能不值得引入。

总而言之，**当我们将硬编码的值与变量进行比较时，编译器必须假定它们具有某种特定的类型，并遵循某些规则来实现这一点**。
如果你发现自己处理的各种类型都可以为 `nil`，那么避免问题的一个常用技巧就是明确地把 `nil` 分配给变量，而不是间接赋值。也就是说，不要写 `a = b`，而写:
```go
var a *int = nil
var b interface{}

if a == nil {
  b = nil
}
```


> 参见：https://www.calhoun.io/when-nil-isnt-equal-to-nil

### 赋值
自增和自减是语句，而不是表达式，因此`x = i++`之类的表达式是错误的。


### 作用域
一个声明语句将程序中的实体和一个名字关联，比如一个函数或一个变量。声明语句的作用域是指源代码中可以有效使用这个名字的范围。

**当编译器遇到一个名字引用时，它会对其定义进行查找，查找过程从最内层的词法域向全局的作用域进行**。如果查找失败，则报告“未声明的名字”这样的错误。**如果该名字在内部和外部的块分别声明过，则内部块的声明首先被找到。在这种情况下，内部声明屏蔽了外部同名的声明，让外部的声明的名字无法被访问**：

```go
func f() {}

var g = "g"

func main() {
    f := "f"
    fmt.Println(f) // "f"; local var f shadows package-level func f
    fmt.Println(g) // "g"; package-level var
    fmt.Println(h) // compile error: undefined: h
}
```

## 基本数据类型
> 之前的章节忘记写了，那就从现在开始吧

### string
* 一个字符串是一个**不可改变**的字节序列。
* 内置的`len`函数可以返回一个字符串中的**字节**数目（不是rune字符数目）。
* 一个原生的字符串面值形式是 \`...\`，使用反引号代替双引号。在原生的字符串面值中，没有转义操作；全部的内容都是字面的意思，包含退格和换行，因此一个程序中的原生字符串面值可能跨越多行（在原生字符串面值内部是无法直接写 \` 字符的，可以用八进制或十六进制转义或 + "`" 连接字符串常量完成）。
    ```go
    const GoUsage = `Go is a tool for managing Go source code.

    Usage:
        go command [arguments]
    ...`
    ```

* 通用的表示一个Unicode码点的数据类型是int32，也就是Go语言中rune对应的类型；它的同义词rune符文正是这个意思。
* Go语言的`range`循环在处理字符串的时候，会**自动隐式解码**UTF8字符串。（注意观察索引的变化，需要注意的是对于非ASCII，索引更新的步长将超过1个字节）
  ![alt range循环字符串时的示意图](https://raw.githubusercontent.com/zoooooway/picgo/master/ch3-05.png)
* 当向`bytes.Buffer`添加任意字符的UTF8编码时，最好使用`bytes.Buffer`的`WriteRune`方法，但是`WriteByte`方法对于写入类似 '[' 和 ']' 等ASCII字符则会更加有效。

### 常量
* 常量表达式的值在编译期计算，而不是在运行期。
* 每种常量的潜在类型都是基础类型：boolean、string或数字。（这意味这常量不能声明引用数据类型）
* 许多常量并没有一个明确的基础类型，编译器为这些没有明确基础类型的数字常量提供比基础类型更高精度的算术运算；你可以认为至少有256bit的运算精度。这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。
* **只有常量可以是无类型的。**
* 不同写法的常量除法表达式可能对应不同的结果：
    ```go
    var f float64 = 212
    fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
    fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
    fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float
    ```
* 当一个无类型的常量被赋值给一个变量的时候，就像下面的第一行语句，或者出现在有明确类型的变量声明的右边，如下面的其余三行语句，无类型的常量将会被隐式转换为对应的类型，如果转换合法的话。

    ```go
    var f float64 = 3 + 0i // untyped complex -> float64
    f = 2                  // untyped integer -> float64
    f = 1e123              // untyped floating-point -> float64
    f = 'a'                // untyped rune -> float64
    ```
    上面的语句相当于:
    ```go
        var f float64 = float64(3 + 0i)
        f = float64(2)
        f = float64(1e123)
        f = float64('a')
    ```
* **无类型整数常量转换为int，它的内存大小是不确定的**，但是无类型浮点数和复数常量则转换为内存大小明确的float64和complex128。
#### 常量生成器iota
在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。

```go
// e.g.
const (
    _ = 1 << (10 * iota)
    KiB // 1024
    MiB // 1048576
    GiB // 1073741824
    TiB // 1099511627776             (exceeds 1 << 32)
    PiB // 1125899906842624
    EiB // 1152921504606846976
    ZiB // 1180591620717411303424    (exceeds 1 << 64)
    YiB // 1208925819614629174706176
)
```


## 复合数据类型
复合数据类型是以不同的方式组合基本类型而构造出来的。

数组和结构体是聚合类型，它们的值由许多元素或成员字段的值组成。数组是由同构的元素组成——**每个数组元素都是完全相同的类型**——结构体则是由异构的元素组成的。数组和结构体都是有固定内存大小的数据结构。相比之下，slice和map则是动态的数据结构，它们将根据需要动态增长。

### 数组
数组是一个由**固定长度**的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。
数组的每个元素可以通过索引下标来访问，索引下标的范围是从0开始到数组长度减1的位置。内置的len函数将返回数组中元素的个数。

可通过字面量形式初始化数组。默认情况下，数组的每个元素都被初始化为元素类型对应的零值。
```go
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) // "0"
// 如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。
c := [...]int{1, 2, 3}
fmt.Printf("%T\n", q) // "[3]int"

// 没有用到的索引可以省略，下面表示索引为99处的元素为-1，因此其他的元素为默认零值
r := [...]int{99: -1}
```

### Slice
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。
一个slice是一个轻量级的数据结构，提供了访问数组子序列（或者全部）元素的功能，而且slice的底层确实引用一个数组对象。

**一个slice由三个部分构成：指针、长度和容量**。**指针指向第一个slice元素对应的底层数组元素的地址**，要注意的是slice的第一个元素并不一定就是数组的第一个元素。长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。

![alt slice和数组](https://raw.githubusercontent.com/zoooooway/picgo/master/ch4-01.png)

slice的切片操作`s[i:j]`，其中0 ≤ i≤ j≤ cap(s)，用于创建一个新的slice，引用s的从**第i个元素开始到第j-1个元素的子序列**。 新的slice将只有j-i个元素。如果i位置的索引被省略的话将使用0代替，如果j位置的索引被省略的话将使用`len(s)`代替。

**和数组不同的是，slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有全部相等元素**。

一个零值的slice等于`nil`。一个`nil`值的slice并没有底层数组。一个`nil`值的slice的长度和容量都是0，但是也有非`nil`值的slice的长度和容量也是0的，例如`[]int{}`或`make([]int, 3)[3:]`。
与任意类型的`nil`值一样，我们可以**用`[]int(nil)`类型转换表达式来生成一个对应类型slice的`nil`值**。
```go
var s []int    // len(s) == 0, s == nil
s = nil        // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{}    // len(s) == 0, s != nil
```

**如果你需要测试一个slice是否是空的，使用`len(s) == 0`来判断，而不应该用`s == nil`来判断**。除了和`nil`相等比较外，一个`nil`值的slice的行为和其它任意0长度的slice一样；例如reverse(`nil`)也是安全的。除了文档已经明确说明的地方，所有的Go语言函数应该以相同的方式对待`nil`值的slice和0长度的slice。

内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分可以省略，在这种情况下，容量将等于长度。
```go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```
**在底层，make创建了一个匿名的数组变量**，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量。在第一种语句中，slice是整个数组的view。在第二个语句中，slice只引用了底层数组的前len个元素，但是容量将包含整个的数组。额外的元素是留给未来的增长用的。

#### append函数
内置的append函数用于向slice追加元素。

**通常我们并不知道append调用是否导致了内存的重新分配，因此我们也不能确认新的slice和原始的slice是否引用的是相同的底层数组空间**。同样，我们不能确认在原先的slice上的操作是否会影响到新的slice。因此，通常是将append返回的结果直接赋值给输入的slice变量：
```go
runes = append(runes, r)
```
> 对于append函数，当slice所对应的底层数组不足以容纳新元素时，内存可能重新分配(为slice创建新的底层数组空间)，但我们不能确认此操作是否发生。

### Map
在Go语言中，一个map就是一个哈希表的引用。它是一个**无序**的key/value对的集合，其中所有的key都是不同的，然后通过给定的key可以在**常数时间复杂度**内检索、更新或删除对应的value。

map中所有的key都有相同的类型，所有的value也有着相同的类型，但是key和value之间可以是不同的数据类型。其中K对应的**key必须是支持==比较运算符的数据类型**，所以map可以通过测试key是否相等来判断是否已经存在。

内置的make函数可以创建一个map：
```go
ages := make(map[string]int) // mapping from strings to ints
```

也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value：
```go
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
empty := map[string]int{} // 创建一个空map
```
使用内置的delete函数可以删除元素：
```go
delete(ages, "alice") // remove element ages["alice"]
```
所有这些操作是安全的，即使这些元素不在map中也没有关系；如果查找失败将返回value类型对应的**零值**。
有时候可能需要知道对应的元素是否真的是在map之中， 比如如果元素类型是一个数字，你可能需要区分一个已经存在的0，和不存在而返回零值的0，可以像下面这样测试：
```go
age, ok := ages["bob"]
if !ok { /* "bob" is not a key in this map; age == 0. */ }
```
可以结合使用来使代码更加简洁
```go
if age, ok := ages["bob"]; !ok { 
    /* ... */ 
}
```

map上的大部分操作，包括查找、删除、`len` 和 `range` 循环都可以安全工作在 `nil` 值的map上，它们的行为和一个空的map类似。但是向一个 `nil` 值的map存入元素将导致一个`panic`异常。
应此，需要谨记：**在向map存数据前必须先创建map**。

map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作：
```go
_ = &ages["bob"] // compile error: cannot take address of map element
```
禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。
> 不同的是，slice同样具有扩容特性，但对slice的元素却可以直接取址。这是由于slice的实现借助数组，而扩容时，新数组的元素即使内存地址改变，原有数组依然存在。因此扩容前取址的值仍然有效。
> 同理，不能直接对slice取址，例如：
> ```go
> arr := [...]int{1, 2, 3, 4, 5}
> s1 := &(arr[:3]) // invalid operation: cannot take address of (arr[:3]) (value of type []int)
> s := arr[:3]
> s2 := &s // compile success
> s3 := &([...]int{1, 2, 3, 4, 5}) // compile success because array is fixed space
> ```
> 参见: https://stackoverflow.com/a/32496031/17180282

和slice一样，map之间也不能进行相等比较；唯一的例外是和 `nil` 进行比较。
```go
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
empty := map[string]int(nil)
zero := map[string]int{}

fmt.Println(ages == nil)
fmt.Println(empty == nil)
fmt.Println(zero == nil)

fmt.Println(ages == empty) // invalid operation: cannot compare ages == empty (map can only be compared to nil)
empty = nil
fmt.Println(ages == empty) // invalid operation: cannot compare ages == empty (map can only be compared to nil)
```

### Struct
结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。

**如果结构体成员名字是以大写字母开头的，那么该成员就是导出的**；这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员。

一个命名为S的结构体类型将不能再包含S类型的成员：因为**一个聚合的值不能包含它自身**。（该限制同样适用于数组。）但是S类型的结构体可以包含 `*S` 指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等。
结构体类型的零值是每个成员都是零值。通常会将零值作为最合理的默认值。

结构体值也可以用结构体字面值表示，结构体字面值可以指定每个成员的值。
```go
type Point struct{ X, Y int }

p1 := Point{1, 2}

p2 := Point{X:1} // 如果成员被忽略的话将默认用零值
```

因为**在Go语言中，所有的函数参数都是值拷贝传入的**，函数参数将不再是函数调用时的原始变量。**因此，如果要在函数内部修改结构体成员的话，必须使用指针传入结构体**。

### JSON
`JavaScript` 对象表示法（`JSON`）是一种用于发送和接收结构化信息的标准协议。

有如下结构体：
```go
type Movie struct {
    Title  string
    Year   int  `json:"released"`
    Color  bool `json:"color,omitempty"`
    Actors []string
}

var movies = []Movie{
    {Title: "Casablanca", Year: 1942, Color: false,
        Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
    {Title: "Cool Hand Luke", Year: 1967, Color: true,
        Actors: []string{"Paul Newman"}},
    {Title: "Bullitt", Year: 1968, Color: true,
        Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
    // ...
}
```
> 在上面的结构体声明中，Year和Color成员后面的字符串面值是结构体成员**Tag**。

将一个Go语言中的结构体 `slice` 转为 `JSON` 的过程叫**编组**（marshaling）。编组通过调用 `json.Marshal` 函数完成, 该函数将返还一个编码后的字节 `slice`，包含很长的字符串，并且没有空白缩进：
```go
data, err := json.Marshal(movies)
if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
```
`json.MarshalIndent` 函数将产生整齐缩进的输出。该函数有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进：
```go
data, err := json.MarshalIndent(movies, "", "    ")
if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
/*
[
    {
        "Title": "Casablanca",
        "released": 1942,
        "Actors": [
            "Humphrey Bogart",
            "Ingrid Bergman"
        ]
    },
    {
        "Title": "Cool Hand Luke",
        "released": 1967,
        "color": true,
        "Actors": [
            "Paul Newman"
        ]
    },
    {
        "Title": "Bullitt",
        "released": 1968,
        "color": true,
        "Actors": [
            "Steve McQueen",
            "Jacqueline Bisset"
        ]
    }
]
*/
```
在编码时，默认使用Go语言结构体的成员名字作为JSON的对象。并且**只有导出的结构体成员才会被编码**。

一个结构体成员Tag是和在编译阶段关联到该成员的元信息字符串：
```go
Year  int  `json:"released"`
Color bool `json:"color,omitempty"`
```
结构体的成员Tag可以是任意的字符串面值，但是通常是一系列用空格分隔的key:"value"键值对序列；因为值中含有双引号字符，因此成员Tag一般用原生字符串面值的形式书写。

json开头键名对应的值用于控制encoding/json包的编码和解码的行为，并且encoding/...下面其它的包也遵循这个约定。

成员Tag中json对应值的第一部分用于指定JSON对象的名字，比如将Go语言中的TotalCount成员对应到JSON中的total_count对象。
Color成员的Tag还带了一个额外的omitempty选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）。

编码的逆操作是解码，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫 `unmarshaling`，通过 `json.Unmarshal` 函数完成。
通过定义合适的Go语言数据结构，我们可以选择性地解码 `JSON` 中感兴趣的成员。
如下，定义一个仅包含Title成员的结构体，解码后的结构体数组就仅包含Title信息。
```go
var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
```

### 文本和HTML模板
由 *text/template* 和 *html/template* 等模板包提供一个将变量值填充到一个文本或HTML格式的模板的机制。
示例：
```go
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
```
在模板中使用的函数分为内置函数和注册函数，内置函数由GO语言内置实现，在任何模板中都可以使用。如果内置函数并没有提供能满足我们需要的函数，我们可以自己实现，并将其注册到该模板中：
```go
var report = template.Must(template.New("issuelist").
    Funcs(template.FuncMap{"daysAgo": daysAgo}). // 注册自定义函数
    Parse(templ))
```
> template.Must辅助函数简化模板编译时的错误处理：它接受一个模板和一个 `error` 类型的参数，检测 `error` 是否为`nil`（如果不是`nil`则发出`panic`异常），然后返回传入的模板。

> 上述为text/template包用法，与其类似，html/template 也提供相同机制，并且其附带html字符转义。

## 函数
函数可以让我们将一个语句序列打包为一个单元，然后可以从程序中其它地方多次调用。

### 函数声明
```go
func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(3,4)) // "5"
```
上述，x和y是形参名，3和4是调用时的传入的实参，函数返回了一个`float64`类型的值。 
返回值也可以像形式参数一样被**命名**（有名）。也就是说，如果函数声明的返回值是命名的，那么，**每个有名返回值会被声明成一个局部变量，并根据该返回值的类型，将其初始化为零值**。 如果一个函数在声明时，包含返回值列表，该函数必须以`return` 语句结尾，除非函数明显无法运行到结尾处。

**函数的类型被称为函数的标识符**。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型或标识符。形参和返回值的变量名不影响函数标识符，也不影响它们是否可以以省略参数类型的形式表示。

**在函数体中，函数的形参作为局部变量，被初始化为调用者提供的值。函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。**
> 在作用域一节中提到，“当编译器遇到一个名字引用时，它会对其定义进行查找，查找过程从最内层的词法域向全局的作用域进行”，并且内层声明将屏蔽外层声明。因此，在函数外声明的同名全局变量将被同名形参和返回值屏蔽。并且，不能在函数中声明同名变量，即使它们的类型不同。

**实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。**

能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数标识符。
```go
package math

func Sin(x float64) float //implemented in assembly language
```

如果一个函数所有的返回值都有显式的变量名，那么该函数的return语句可以省略操作数。这称之为bare return。
```go
// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}
func countWordsAndImages(n *html.Node) (words, images int) { /* ... */ }
```
按照返回值列表的次序，返回所有的返回值，在上面的例子中，每一个return语句等价于：
```go
return words, images, err
```
> 虽然良好的命名很重要，但你也不必为每一个返回值都取一个适当的名字。比如，按照惯例，函数的最后一个bool类型的返回值表示函数是否运行成功，error类型的返回值代表函数的错误信息，对于这些类似的惯例，我们不必思考合适的命名，它们都无需解释。

当一个函数有多处return语句以及许多返回值时，bare return 可以减少代码的重复，**但是使得代码难以被理解**。

### error
**对于大部分函数而言，永远无法确保能否成功运行。**

对于那些将运行失败看作是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个，来传递错误信息。如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为ok。比如，cache.Lookup失败的唯一原因是key不存在，那么代码可以按照下面的方式组织：
```go
value, ok := cache.Lookup(key)
if !ok {
    // ...cache[key] does not exist…
}
```

对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。

### 函数值
在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。
对函数值（function value）的调用类似函数调用：
```go
func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

f := square
fmt.Println(f(3)) // "9"

f = negative // 可以相互赋值为相同描述符的函数
fmt.Println(f(3))     // "-3"
fmt.Printf("%T\n", f) // "func(int) int"

f = product // compile error: can't assign func(int, int) int to func(int) int
```

**函数类型的零值是`nil`。调用值为`nil`的函数值会引起`panic`错误**：
```go
var f func(int) int
f(3) // 此处f的值为nil, 会引起panic错误
```

**函数值可以与`nil`比较**：
```go
var f func(int) int
if f != nil {
    f(3)
}
```
但是**函数值之间是不可比较的，也不能用函数值作为`map`的`key`**。

函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为(函数)。
```go
// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    if post != nil {
        post(n)
    }
}
```
### 匿名函数
拥有函数名的函数只能在包级语法块中被声明，通过函数字面量（function literal），我们可绕过这一限制，在任何表达式中表示一个函数值。函数字面量的语法和函数声明相似，区别在于func关键字后没有函数名。**函数值字面量是一种表达式，它的 值 被称为匿名函数（anonymous function）。**

函数字面量允许我们在使用函数时，再定义它。
更为重要的是，通过这种方式定义的函数可以访问完整的词法环境（lexical environment），这意味着**在函数中定义的内部函数可以引用该函数的变量**：
```go
// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f1 := squares()
	f2 := squares()
	fmt.Println(f1()) // "1"
	fmt.Println(f1()) // "4"
	fmt.Println(f1()) // "9"
	fmt.Println("---------------------")
	fmt.Println(f2()) // "1"
	fmt.Println(f2()) // "4"
	fmt.Println(f2()) // "9"
}
```
函数`squares`返回另一个类型为 `func() int` 的函数。对`squares`的一次调用会**生成一个局部变量** `x`并返回一个匿名函数。每次调用匿名函数时，该函数都会先使`x`的值加1，再返回`x`的平方。第二次调用`squares`时，会生成第二个`x`变量，并返回一个新的匿名函数。新匿名函数操作的是第二个x变量。

**函数值不仅仅是一串代码，还记录了状态**。在`squares`中定义的匿名内部函数可以访问和更新`squares`中的局部变量，这意味着匿名函数和`squares`中，存在**变量引用**。这就是函数值属于**引用类型**和函数值**不可比较**的原因。
`Go`使用闭包（closures）技术实现函数值，`Go`程序员也把函数值叫做闭包。
>通过这个例子，我们看到变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于f中。因此，变量x的生命周期现在由f决定了。

#### 警告：捕获迭代变量
`Go`使用闭包（closures）技术实现函数值，因此，在循环中的函数值引用迭代变量时非常容易出错。
正确做法：
```go
var rmdirs []func()
for _, d := range tempDirs() {
    dir := d // NOTE: necessary!
    os.MkdirAll(dir, 0755) // creates parent directories too
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)
    })
}
// ...do some work…
for _, rmdir := range rmdirs {
    rmdir() // clean up
}
```
错误做法：
```go
var rmdirs []func()
for _, dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect!
    })
}
```
如果不将dir变量在循环中赋给另一个变量来保存，那么在函数值中，引用的变量dir将会是最后一次循环完成之后的dir值。也就是说，每一个函数值中的变量dir都是相同的，而不是预期的每次迭代的值。
> https://docs.hacknode.org/gopl-zh/ch5/ch5-06.html


### 可变参数
参数数量可变的函数称为可变参数函数。
在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号“...”，这表示该函数会接收任意数量的该类型参数。
**在函数体中,可变参数被看作是切片类型的变量。**

虽然在可变参数函数内部，...int 型参数的行为看起来很像切片类型，但实际上，**可变参数函数和以切片作为参数的函数是不同的**。
```go
func f(...int) {}
func g([]int) {}
fmt.Printf("%T\n", f) // "func(...int)"
fmt.Printf("%T\n", g) // "func([]int)"
```

### Deferred函数
`defer`语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。
在调用普通函数或方法前加上关键字`defer`，就完成了`defer`所需要的语法。

当**执行到**关键字`defer`所在语句时，函数和参数表达式得到**计算**，但直到包含该`defer`语句的函数**执行完毕**时，`defer`后的函数才会被**执行**，不论包含`defer`语句的函数是通过`return`正常结束，还是由于`panic`导致的异常结束。
```go
func ReadFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return ReadAll(f)
}
```
> 释放资源的defer应该直接跟在请求资源的语句后。

对匿名函数采用defer机制，可以使其观察函数的返回值。
```go
func double(x int) (result int) {
    defer func() { fmt.Printf("double(%d) = %d\n", x,result) }()
    return x + x
}
_ = double(4)
// Output:
// "double(4) = 8"
```

被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值：
```go
func triple(x int) (result int) {
    defer func() { result += x }()
    return double(x)
}
fmt.Println(triple(4)) // "12"
```

**在循环体中的defer语句需要特别注意**，因为只有在函数执行完毕后，这些被延迟的函数才会执行。下面的代码会导致系统的文件描述符耗尽，因为在所有文件都被处理之前，没有文件会被关闭。
```go
for _, filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close() // NOTE: risky; could run out of file descriptors
    // ...process f…
}
```

一种解决方法是将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数。
```go
for _, filename := range filenames {
    if err := doFile(filename); err != nil {
        return err
    }
}
func doFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    // ...process f…
}
```

### Panic异常
Go的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起`painc`异常。

一般而言，当`panic`异常发生时，程序会中断运行，并立即执行在该`goroutine`中被延迟的函数（`defer`机制）。

不是所有的`panic`异常都来自运行时，直接调用内置的`panic`函数也会引发`panic`异常；`panic`函数接受任何值作为参数。当某些不应该发生的场景发生时，我们就应该调用`panic`。

**在Go的panic机制中，延迟函数的调用在释放堆栈信息之前**。这意味着，即使程序因为`panic`异常而将要退出，也可以在延迟函数中访问堆栈信息。
```go
func main() {
    defer printStack()
    f(3)
}
func printStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], false)
    os.Stdout.Write(buf[:n])
}
```

### Recover捕获异常
通常来说，不应该对`panic`异常做任何处理，但有时，也许我们可以从异常中恢复，至少我们可以在程序崩溃前，做一些操作。

如果在`deferred`函数中调用了内置函数`recover`，并且定义该`defer`语句的函数发生了`panic`异常，`recover`会使程序从`panic`中恢复，并返回`panic value`。导致`panic`异常的函数不会继续运行，但能**正常返回**。**在未发生panic时调用`recover`，`recover`会返回`nil`。** 
```go
func f() (err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error: %v", p)
        }
    }()
    // ...parser...
}
```

不加区分的恢复所有的`panic`异常，不是可取的做法；因为**在`panic`之后，无法保证包级变量的状态仍然和我们预期一致**。比如，对数据结构的一次重要更新没有被完整完成、文件或者网络连接没有被关闭、获得的锁没有被释放。此外，如果写日志时产生的`panic`被不加区分的恢复，可能会导致漏洞被忽略。

**作为被广泛遵守的规范，你不应该试图去恢复其他包引起的`panic`**。 公有的`API`应该将函数的运行失败作为`error`返回，而不是`panic`。同样的，你也不应该恢复一个由他人开发的函数引起的`panic`，比如说调用者传入的回调函数，因为你无法确保这样做是安全的。

有时我们很难完全遵循规范，举个例子，`net/http`包中提供了一个`web`服务器，将收到的请求分发给用户提供的处理函数。很显然，我们不能因为某个处理函数引发的`panic`异常，杀掉整个进程；`web`服务器遇到处理函数导致的`panic`时会调用`recover`，输出堆栈信息，继续运行。这样的做法在实践中很便捷，但也会引起资源泄漏，或是因为`recover`操作，导致其他问题。

**基于以上原因，安全的做法是有选择性的`recover`**。换句话说，**只恢复应该被恢复的panic异常**，此外，这些异常所占的比例应该尽可能的低。


## 方法
方法是基于面向对象编程（OOP）的概念。

一个对象其实也就是一个简单的值或者一个变量，在这个对象中会包含一些方法，而**一个方法则是一个一个和特殊类型关联的函数**。一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象，而是借助方法来做这些事情。

### 方法声明
在函数声明时，在其名字之前放上一个变量，这就是一个方法的声明方式。
这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个**独占**的方法。
```go
import "math"

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) X(q Point) {} // error: field and method with the same name X
```
上面的代码中，参数p，叫做方法的接收器（receiver）。
`p.Distance`的表达式叫做选择器，因为他会选择合适的对应p这个对象的`Distance`方法来执行。选择器也会被用来选择一个`struct`类型的字段，比如`p.X`。
由于方法和字段都是在同一命名空间，所以**对象的字段和对象的方法名重复的话，则会编译报错**。

让我们来定义一个Path类型，这个Path代表一个线段的集合，并且也给这个Path定义一个叫Distance的方法。
```go
// A Path is a journey connecting the points with straight lines.
type Path []Point
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i-1].Distance(path[i])
        }
    }
    return sum
}
```
Path是一个命名的slice类型，而不是Point那样的struct类型，然而我们依然可以为它定义方法。

在Go语言里，我们为一些简单的数值、字符串、slice、map来定义一些附加行为很方便。我们可以给同一个包内的**任意命名类型**定义方法，只要这个命名类型的底层类型（译注：这个例子里，底层类型是指[]Point的底层类型是slice，Path就是命名类型）不是指针或者interface。

**对于一个给定的类型，其内部的方法都必须有唯一的方法名，但是不同的类型却可以有同样的方法名。**

### 基于指针对象的方法
**当调用一个函数时，会对其每一个参数值进行拷贝**，如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，这种情况下我们就需要用到指针了。

例如用来更新接收器的对象的方法，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法，如下：
```go
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}
```
这个方法的名字是`(*Point).ScaleBy`。这里的括号是必须的；没有括号的话这个表达式可能会被理解为`*(Point.ScaleBy)`。

在现实的程序里，一般会约定如果Point这个类有一个指针作为接收器的方法，那么所有Point的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数。

只有类型（Point）和指向他们的指针(*Point)，才可能是出现在接收器声明里的两种接收器。此外，为了避免歧义，在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的，比如下面这个例子：
```go
type P *int
func (P) f() { /* ... */ } // compile error: invalid receiver type
```

不管你的`method`的`receiver`是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。

如果命名类型`T`（译注：用`type xxx`定义的类型）的所有方法都是用`T`类型自己来做接收器（而不是`*T`），那么拷贝这种类型的实例就是安全的；调用他的任何一个方法也就会产生一个值的拷贝。比如`time.Duration`的这个类型，在调用其方法时就会被全部拷贝一份，包括在作为参数传入函数的时候。
> 与函数调用会拷贝参数的值是相同道理，调用方法会拷贝接收器的值。

但是如果一个方法使用指针作为接收器，你需要避免对其进行拷贝，因为这样可能会破坏掉该类型内部的不变性。比如你对`bytes.Buffer`对象进行了拷贝，那么可能会引起原始对象和拷贝对象只是别名而已，实际上它们指向的对象是一样的。紧接着对拷贝后的变量进行修改可能会有让你有意外的结果。
> 以指针作为接收器来调用方法，如果拷贝这个指针，并且通过方法将其传输到了其他地方进行持有。那么，该指针就可能被修改并且影响原始对象的值，而且我们完全不知道会不会被修改。

#### Nil也是一个合法的接收器类型
就像一些函数允许nil指针作为参数一样，方法理论上也可以用nil指针作为其接收器，尤其当nil对于对象来说是合法的零值时，比如map或者slice。

因为nil的字面量编译器无法判断其准确类型，所以想要使用`nil`来作为接收器调用方法时，需要先转换为对应类型的`nil`值。
```go
c := (*Point)(nil)
c.Distance(Point{1, 2}) // compile success

nil.Distance(Point{1, 2}) // compile error: nil.Distance undefined (type untyped nil has no field or method Distance)

(*Point)(nil).Distance(Point{1, 2}) // compile success
```

### 通过嵌入结构体来扩展类型
结构体的内嵌可以使我们在定义结构体时得到一种句法上的简写形式，并使其包含内嵌结构体类型所具有的一切字段，然后再定义一些自己的。如果我们想要的话，我们可以直接认为通过嵌入的字段就是结构体自身的字段，而完全不需要在调用时指出内嵌的结构体是什么。
同理，对于结构体的方法也是一样。我们可以把一个结构体类型A当作接收器来调用该结构体里内嵌机构体B所具有的方法，即使结构体A里没有声明这些方法。

在类型中内嵌的匿名字段也可能是一个命名类型的指针，这种情况下字段和方法会被间接地引入到当前的类型中（译注：访问需要通过该指针指向的对象去取）。添加这一层间接关系让我们可以共享通用的结构并动态地改变对象之间的关系。下面这个ColoredPoint的声明内嵌了一个`*Point`的指针。
```go
type ColoredPoint struct {
    *Point
    Color color.RGBA
}

p := ColoredPoint{&Point{1, 1}, red}
q := ColoredPoint{&Point{5, 4}, blue}
fmt.Println(p.Distance(*q.Point)) // "5"
q.Point = p.Point                 // p and q now share the same Point
p.ScaleBy(2)
fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
```

虽然方法只能在命名类型（像Point）或者指向类型的指针上定义。但是通过内嵌，有些时候我们也可以在匿名`struct`类型上定义方法。
```go
var cache = struct {
    sync.Mutex
    mapping map[string]string
}{
    mapping: make(map[string]string),
}


func Lookup(key string) string {
    cache.Lock()
    v := cache.mapping[key]
    cache.Unlock()
    return v
}
```
匿名结构体cache调用的方法其实是内部的内嵌结构体Mutex类型的方法，但在我们看来，这就像他自己定义的方法一样。

### 方法值和方法表达式

我们经常选择一个方法，并且在同一个表达式里执行，比如常见的p.Distance()形式，实际上将其分成两步来执行也是可能的。
`p.Distance`叫作“选择器”，选择器会返回一个方法“值”->一个将方法（`Point.Distance`）绑定到特定接收器变量的函数。这个函数可以不通过指定其接收器即可被调用；即调用时不需要指定接收器(因为已经在前文中指定过了），只要传入函数的参数即可。

在一个包的API需要一个函数值、且调用方希望操作的是某一个绑定了对象的方法的话，方法“值”会非常实用。

举例来说，下面例子中的`time.AfterFunc`这个函数的功能是在指定的延迟时间之后来执行一个（译注：另外的）函数。且这个函数操作的是一个`Rocket`对象`r`
```go
type Rocket struct { /* ... */ }
func (r *Rocket) Launch() { /* ... */ }
r := new(Rocket)

time.AfterFunc(10 * time.Second, func() { r.Launch() })
```
直接用方法“值”传入AfterFunc的话可以更为简短：
```go
time.AfterFunc(10 * time.Second, r.Launch)
```
> 省去了定义匿名函数

和方法“值”相关的还有**方法表达式**。当调用一个方法时，与调用一个普通的函数相比，我们必须要用选择器（`p.Distance`）语法来指定方法的接收器。

当`T`是一个类型时，方法表达式可能会写作`T.f`或者`(*T).f`，会返回一个函数“值”，这种函数会将其第一个参数用作接收器，所以可以用通常（译注：不写选择器）的方式来对其进行调用：
```go
p := Point{1, 2}
q := Point{4, 6}

distance := Point.Distance   // method expression
fmt.Println(distance(p, q))  // "5"
fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

scale := (*Point).ScaleBy
scale(&p, 2)
fmt.Println(p)            // "{2 4}"
fmt.Printf("%T\n", scale) // "func(*Point, float64)"

// 译注：这个Distance实际上是指定了Point对象为接收器的一个方法func (p Point) Distance()，
// 但通过Point.Distance得到的函数需要比实际的Distance方法多一个参数，
// 即其需要用第一个额外参数指定接收器，后面排列Distance方法的参数。
// 看起来本书中函数和方法的区别是指有没有接收器，而不像其他语言那样是指有没有返回值。
```

当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了。你可以根据选择来调用接收器各不相同的方法。下面的例子，变量`op`代表`Point`类型的`addition`或者`subtraction`方法，`Path.TranslateBy`方法会为其`Path`数组中的每一个`Point`来调用对应的方法：
```go
type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
    var op func(p, q Point) Point
    if add {
        op = Point.Add
    } else {
        op = Point.Sub
    }
    for i := range path {
        // Call either path[i].Add(offset) or path[i].Sub(offset).
        path[i] = op(path[i], offset)
    }
}
```

#### 示例: Bit数组
```go
// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
    words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
    word, bit := x/64, uint(x%64)
    return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
    word, bit := x/64, uint(x%64)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < 64; j++ {
            if word&(1<<uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", 64*i+j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}
```

当你为一个复杂的类型定义了一个`String`方法时，`fmt`包就会特殊对待这种类型的值，这样可以让这些类型在打印的时候看起来更加友好，而不是直接打印其原始的值。`fmt`会直接调用用户定义的`String`方法。这种机制依赖于接口和类型断言。

这里要注意：我们声明的`String`和`Has`两个方法都是以指针类型`*IntSet`来作为接收器的，但实际上对于这两个类型来说，把接收器声明为指针类型也没什么必要。不过另外两个函数就不是这样了，因为另外两个函数操作的是`s.words`对象，如果你不把接收器声明为指针对象，那么实际操作的是拷贝对象，而不是原来的那个对象。因此，因为我们的`String`方法定义在`IntSet`指针上，所以当我们的变量是`IntSet`类型而不是`IntSet`指针时，可能会有下面这样让人意外的情况：
```go
fmt.Println(&x)         // "{1 9 42 144}"
fmt.Println(x.String()) // "{1 9 42 144}"
fmt.Println(x)          // "{[4398046511618 0 65536]}"
```

在第一个`Println`中，我们打印一个`*IntSet`的指针，这个类型的指针确实有自定义的`String`方法。第二`Println`，我们直接调用了`x`变量的`String()`方法；这种情况下编译器会隐式地在`x`前插入`&`操作符（前面提到过，编译器在这种情况下会帮我们做隐式插入），这样相当于我们还是调用的`IntSet`指针的`String`方法。在第三个`Println`中，因为`IntSet`类型没有`String`方法，所以`Println`方法会直接以原始的方式理解并打印。所以在这种情况下`&`符号是不能忘的。

### 封装
`Go`语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写字母的则不会。因而如果我们想要封装一个对象，我们必须将其定义为一个`struct`。

这种基于名字的手段使得在语言中最小的封装单元是`package`，而不是像其它语言一样的类型。一个`struct`类型的字段对同一个包的所有代码都有可见性，无论你的代码是写在一个函数还是一个方法里。

封装提供了三方面的优点。
* 首先，因为**调用方不能直接修改对象的变量值**，其只需要关注少量的语句并且只要弄懂少量变量的可能的值即可。
* 第二，**隐藏实现的细节**，可以**防止调用方依赖那些可能变化的具体实现**，这样使设计包的程序员在不破坏对外的api情况下能得到更大的自由。
* 封装的第三个优点也是最重要的优点，是**阻止了外部调用方对对象内部的值任意地进行修改**。因为对象内部变量只可以被同一个包内的函数修改，所以包的作者可以让这些函数确保对象内部的一些值的不变性。
  > 因为只要不声明导出，外部则无法访问到内部值。

只用来访问或修改内部变量的函数被称为`setter`或者`getter`。
在命名一个`getter`方法时，我们通常会省略掉前面的Get前缀。这种简洁上的偏好也可以推广到各种类型的前缀比如Fetch，Find或者Lookup（与`Java`的命名风格有些不同）。

`Go`的编码风格不禁止直接导出字段。
一旦进行了导出，就没有办法在保证API兼容的情况下去除对其的导出，所以在一开始的选择一定要经过深思熟虑并且要考虑到包内部的一些不变量的保证，未来可能的变化，以及调用方的代码质量是否会因为包的一点修改而变差。

封装并不总是理想的。

将`IntSet`和本章开头的`geometry.Path`进行对比。`Path`被定义为一个`slice`类型，这允许其调用`slice`的字面方法来对其内部的`points`用`range`进行迭代遍历；在这一点上，`IntSet`是没有办法让你这么做的。

这两种类型决定性的不同：`geometry.Path`的本质是一个坐标点的序列，不多也不少，我们可以预见到之后也并不会给他增加额外的字段，所以在`geometry`包中将`Path`暴露为一个`slice`。相比之下，`IntSet`仅仅是在这里用了一个`[]uint64的slice`。这个类型还可以用`[]uint`类型来表示，或者我们甚至可以用其它完全不同的占用更小内存空间的东西来表示这个集合，所以我们可能还会需要额外的字段来在这个类型中记录元素的个数。也正是因为这些原因，我们让`IntSet`对调用方不透明。

## 接口
**接口类型是对其它类型行为的抽象和概括。**
因为**接口类型不会和特定的实现细节绑定在一起**，通过这种抽象的方式我们可以让我们的函数更加灵活和更具有适应能力。

`Go`语言中接口类型的独特之处在于它是满足**隐式实现**的。也就是说，我们没有必要对于给定的具体类型定义所有满足的接口类型；简单地拥有一些必需的方法就足够了。
这种设计可以让你创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义；当我们使用的类型来自于不受我们控制的包时这种设计尤其有用。

### 接口约定
一个具体的类型可以准确的描述它所代表的值，并且展示出对类型本身的一些操作方式：就像数字类型的算术操作，切片类型的取下标、添加元素和范围获取操作。具体的类型还可以通过它的内置方法提供额外的行为操作。总的来说，**当你拿到一个具体的类型时你就知道它的本身是什么和你可以用它来做什么。**

在Go语言中还存在着另外一种类型：接口类型。接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会表现出它们自己的方法。也就是说**当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。**

### 接口类型
接口类型具体描述了一系列方法的集合，**一个实现了这些方法的具体类型是这个接口类型的实例**。

一些新的接口类型通过组合已有的接口来定义。下面是两个例子：
```go
type ReadWriter interface {
    Reader
    Writer
}
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

不使用内嵌来声明io.ReadWriter接口。
```go
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
```

或者甚至使用一种混合的风格：
```go
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Writer
}
```

上面三种方式的效果完全相同，方法顺序的变化也没有影响，唯一重要的就是这个集合里面的方法。

### 实现接口的条件
一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。

表达一个类型属于某个接口只要这个类型实现这个接口。所以：
```go
var w io.Writer
w = os.Stdout           // OK: *os.File has Write method
w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
w = time.Second         // compile error: time.Duration lacks Write method

var rwc io.ReadWriteCloser
rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
```

这个规则甚至适用于等式右边本身也是一个接口类型:
```go
w = rwc                 // OK: io.ReadWriteCloser has Write method
rwc = w                 // compile error: io.Writer lacks Close method
```

关于`interface{}`类型，它没有任何方法，因此`interface{}`被称为空接口类型。
因为空接口类型对实现它的类型没有要求，所以**我们可以将任意一个值赋给空接口类型**。

对于每一个命名过的具体类型`T`；它的一些方法的接收者是类型T本身然而另一些则是一个`*T`的指针。还记得在`T`类型的参数上调用一个`*T`的方法是合法的，只要这个参数是一个变量；编译器隐式的获取了它的地址。但这仅仅是一个语法糖：T类型的值不拥有所有`*T`指针的方法，这样它就可能只实现了更少的接口。
```go
type IntSet struct { /* ... */ }
func (*IntSet) String() string
var _ = IntSet{}.String() // compile error: String requires *IntSet receiver

var s IntSet
var _ = s.String() // OK: s is a variable and &s has a String method

var _ fmt.Stringer = &s // OK
var _ fmt.Stringer = s  // compile error: IntSet lacks String method
```
> 需要注意的就是`IntSet{}`这种字面量表达式不会使编译器去隐式获取表达式的指针。

非空的接口类型比如`io.Writer`经常被指针类型实现，尤其当一个或多个接口方法像Write方法那样隐式的给接收者带来变化的时候。一个结构体的指针是非常常见的承载方法的类型。

### 接口值
**接口值，由两个部分组成，一个具体的类型(type)和那个类型的值(value)。**
> 变量一节提到：每个变量都具备类型和值。假设存在变量a，则可以用`fmt.Printf("a=(%T, %v)\n", a, a)`来查看其类型和值。

它们被称为接口的**动态类型**和**动态值**。
> 动态类型指实现这个接口的类型（比如结构体），动态值就是这个类型的对应值

对于像`Go`语言这种静态类型的语言，类型是编译期的概念，**因此一个类型不是一个值**。
> 不大能理解这句话什么意思...

下面4个语句中，变量w得到了3个不同的值。（声明时的值和最后一行的值是相同的）
```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

在`Go`语言中，变量总是被一个定义明确的值初始化，即使接口类型也不例外。**一个接口的零值就是它的类型和值的部分都是`nil`**。
![](https://raw.githubusercontent.com/zoooooway/picgo/master/202302251758900.png)

**一个接口值基于它的动态类型被描述为空或非空**，所以在`var w io.Writer`这里，`w`是一个空的接口值。
> 因为`w`被声明为`io.Writer`的接口值，而声明语句中没有给出`w`动态类型即`io.Writer`的具体实现类型，因此`w`的值为空

可以通过使用`w==nil`或者`w!=nil`来判断接口值是否为空。调用一个空接口值上的任意方法都会产生`panic`:
```go
w.Write([]byte("hello")) // panic: nil pointer dereference
```

</br>

第二个语句将一个`*os.File`类型的值赋给变量`w`:
```go
w = os.Stdout
```

这个赋值过程调用了一个具体类型到接口类型的隐式转换，这和显式的使用`io.Writer(os.Stdout)`是等价的。
这个接口值的动态类型被设为`*os.File`指针的类型描述符，它的动态值持有`os.Stdout`的拷贝；这是一个代表处理标准输出的`os.File`类型变量的指针。
![](https://raw.githubusercontent.com/zoooooway/picgo/master/202302252025618.png)

通常在编译期，我们不知道接口值的动态类型是什么，所以一个接口上的调用必须使用动态分配。
因为不是直接进行调用，所以编译器必须**把代码生成在类型描述符的`Write`方法上，然后间接调用那个地址**。这个**调用的接收者是一个接口动态值的拷贝**：`os.Stdout`（`os.Stdout`是指针，那如果不是指针而是类型值呢?）。效果和下面这个直接调用一样：
```go
os.Stdout.Write([]byte("hello")) // "hello"
```
</br>

第三个语句给接口值赋了一个`*bytes.Buffer`类型的值:
```go
w = new(bytes.Buffer)
```

现在动态类型是`*bytes.Buffer`并且动态值是一个指向新分配的缓冲区的指针。
![](https://raw.githubusercontent.com/zoooooway/picgo/master/202302252032832.png)

</br>

最后，第四个语句将nil赋给了接口值：
```go
w = nil
```

这个重置将它所有的部分都设为`nil`值，**把变量`w`恢复到和它之前定义时相同的状态**，在图7.1中可以看到。

**一个接口值可以持有任意大的动态值**。
例如，表示时间实例的time.Time类型，这个类型有几个对外不公开的字段。我们从它上面创建一个接口值：
```go
var x interface{} = time.Now()
```

结果可能和图7.4相似。从概念上讲，不论接口值多大，动态值总是可以容下它。（这只是一个概念上的模型；具体的实现可能会非常不同）
![](https://raw.githubusercontent.com/zoooooway/picgo/master/202302252043695.png)

接口值可以使用`==`和`!＝`来进行比较。
因为**接口值是可比较的**，所以它们可以用在`map`的键或者作为`switch`语句的操作数。
**两个接口值相等仅当它们都是`nil`值，或者它们的动态类型相同并且动态值也根据这个动态类型的`==`操作相等**。
> 可以参照变量、指针与`nil`一节中的内容来理解

如果两个接口值的动态类型相同，但是这个动态类型是不可比较的（比如切片），将它们进行比较就会失败并且`panic`:
```go
var x interface{} = []int{1, 2, 3}
fmt.Println(x == x) // panic: comparing uncomparable type []int
```

考虑到这点，接口类型是非常与众不同的。其它类型要么是安全的可比较类型（如基本类型和指针）要么是完全不可比较的类型（如切片，映射类型，和函数），但是**在比较接口值或者包含了接口值的聚合类型时，我们必须要意识到潜在的`panic`**。同样的风险也存在于使用接口作为`map`的键或者`switch`的操作数。只能比较你非常确定它们的动态值是可比较类型的接口值。

#### 警告：一个包含nil指针的接口不是nil接口
**一个不包含任何值的`nil`接口值和一个刚好包含`nil`指针的接口值是不同的**。

思考下面的程序。当`debug`变量设置为`true`时，`main`函数会将`f`函数的输出收集到一个`bytes.Buffer`类型中。
```go
const debug = true

func main() {
    var buf *bytes.Buffer
    if debug {
        buf = new(bytes.Buffer) // enable collection of output
    }
    f(buf) // NOTE: subtly incorrect!
    if debug {
        // ...use buf...
    }
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
    // ...do something...
    if out != nil {
        out.Write([]byte("done!\n"))
    }
}
```
> 参考变量、指针与`nil`一节。并解释上述代码在`debug`值为`false`时会发生什么?


### 类型断言
类型断言是一个使用在接口值上的操作。
在语法上，形如`x.(T)`这种声明被称为断言类型，这里`x`表示一个接口的类型，`T`表示一个类型。一个类型断言检查它操作对象的动态类型(`x`)是否和断言的类型(`T`)匹配。
这里有两种可能:
* 如果断言的类型`T`是一个**具体类型**，类型断言检查`x`的动态类型是否和`T`相同。
  如果这个检查成功了，类型断言的结果是`x`的动态值，当然它的类型是`T`。换句话说，具体类型的类型断言从它的操作对象中获得具体的值。如果检查失败，这个操作会抛出`panic`。例如：
    ```go
    var w io.Writer
    w = os.Stdout
    f := w.(*os.File)      // success: f == os.Stdout
    c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
    ```
* 如果相反地断言的类型`T`是一个**接口类型**，类型断言检查`x`的动态类型是否满足`T`接口的要求（判断`x`是否实现了`T`）。
  在下面的第一个类型断言后，`w`和`rw`都持有`os.Stdout`，因此它们都有一个动态类型`*os.File`，但是变量`w`是一个`io.Writer`类型，只对外公开了文件的`Write`方法，而`rw`变量还公开了它的`Read`方法。
    ```go
    var w io.Writer
    w = os.Stdout
    rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
    w = new(ByteCounter)
    rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
    ```

**如果断言操作的对象(`x`)是一个`nil`接口值，那么这个类型断言必定会失败**。

如果类型断言出现在一个预期有两个结果的赋值操作中，例如如下的定义，这个操作不会在失败的时候发生`panic`，但是替代地返回一个额外的第二个结果，这个结果是一个标识成功与否的布尔值：
```go
var w io.Writer = os.Stdout
f, ok := w.(*os.File)      // success:  ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
```
第一个结果等于**被断言类型**的**零值**，在这个例子中就是一个`nil`的`*bytes.Buffer`类型。第二个结果是标识成功与否的布尔值，通常赋值给一个命名为`ok`的变量。

类型断言的结果常用来指示程序接下来该如何执行。当类型断言的操作对象是一个变量，你有时会看见原来的变量名重用而不是声明一个新的本地变量名，这个重用的变量原来的值会被覆盖（理解：其实是声明了一个同名的新的本地变量，外层原来的`w`不会被改变），如下面这样：
```go
if w, ok := w.(*os.File); ok {
    // ...use w...
}
```

### 类型分支
如果这个`if-else`链对一连串值做类型断言测试。`type switch`（类型分支）可以简化类型断言的`if-else`链。
在最简单的形式中，一个类型分支像普通的`switch`语句一样，它的运算对象是`x.(type)`——它使用了关键词字面量`type`——并且每个`case`有一到多个类型。一个类型分支基于这个接口值的动态类型使一个多路分支有效。这个`nil`的`case`和`if x == nil`匹配，并且这个`default`的`case`和如果其它`case`都不匹配的情况匹配
```go
switch x.(type) {
case nil:       // ...
case int, uint: // ...
case bool:      // ...
case string:    // ...
default:        // ...
}
```
>当一个或多个`case`类型是接口时，`case`的顺序就会变得很重要，因为可能会有两个`case`同时匹配的情况。

### 一些建议
当设计一个新的包时，新手Go程序员总是先创建一套接口，然后再定义一些满足它们的具体类型。这种方式的结果就是有很多的接口，它们中的每一个仅只有一个实现。不要再这么做了。这种接口是不必要的抽象，它们也有一个运行时损耗。**接口只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要。**

当一个接口只被一个单一的具体类型实现时有一个例外，就是由于它的依赖，这个具体类型不能和这个接口存在在一个相同的包中。这种情况下，一个接口是解耦这两个包的一个好方式。

当新的类型出现时，小的接口更容易满足。对于接口设计的一个好的标准就是 *ask only for what you need*（只考虑你需要的东西）。

## Goroutines和Channels
`goroutine`和`channel`是Go语言并发支持的核心，其支持“顺序通信进程”（communicating sequential processes）或被简称为CSP。CSP是一种现代的并发编程模型，在这种编程模型中值会在不同的运行实例（`goroutine`）中传递，尽管大多数情况下仍然是被限制在单一实例中。

### Goroutines
在Go语言中，每一个并发的执行单元叫作一个`goroutine`。

当一个程序启动时，其主函数即在一个单独的`goroutine`中运行，我们叫它`main goroutin`e。新的`goroutine`会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的`goroutine`中运行。而go语句本身会迅速地完成。
```go
f()    // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

主函数返回时，所有的`goroutine`都会被直接打断，程序退出。
除了从主函数退出或者直接终止程序之外，没有其它的编程方法能够让一个`goroutine`来打断另一个的执行，但是之后可以看到一种方式来实现这个目的，通过`goroutine`之间的通信来让一个`goroutine`请求其它的`goroutine`，并让被请求的`goroutine`自行结束执行。













