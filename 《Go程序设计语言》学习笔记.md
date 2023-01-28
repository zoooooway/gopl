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