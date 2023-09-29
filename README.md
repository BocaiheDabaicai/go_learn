# Golang

---

## 第一章 起步

#### 1.0 前置知识

1. 使用go命令运行项目

2. 包意味着项目，即工作空间

3. `import xxx`意味着导入某个包，使用到当前的文件上

4. 函数声明的格式

5. 函数组织结构，声明包名、导入使用的包、声明函数

#### 1.1 go命令

| 名称           | 作用              |
| ------------ | --------------- |
| `go build`   | 编译.go源代码文件      |
| `go run`     | 编译并执行一两个.go文件   |
| `go fmt`     | 格式化当前目录下的所有文件代码 |
| `go install` | 编译并升级一个包        |
| `go get`     | 下载一些包的源代码       |
| `go test`    | 运行所有有关当前项目的测试   |

#### 1.2 包名命

代码的首行，例如`pakeage main`，表示当前文件属于包`main`

包分为两种类型

1. `执行型`，进行运行的文件

2. `可重复使用型`，文件的代码可以反复使用

#### 1.3 导入

可导入的包分为两种

1. 标准库

2. 自定义包

#### 1.4 函数声明

四个部分组成，函数声明、函数名、函数形参、函数体

`func name(){}`

#### 1.5 Go简介

`Go`是一门静态类型语言，不像`Java`等语言面向对象

---

## 第二章 卡片（基础知识）

#### 2.1 变量声明

声明变量：变量的存在形式、变量名、变量类型，赋予符合变量类型的数值

示例：

`var card string = "Ace of Spades"`

`card := "Ace of Spades"  // 简写方式，只需声明一次`

#### 2.2 变量的基本类型

| 类型名称      | 类型说明 |
| --------- | ---- |
| `bool`    |      |
| `string`  |      |
| `int`     |      |
| `float64` |      |
| `array`   |      |
| `map`     | 集合   |

#### 2.3 函数实现

上面提到了如何声明一个函数，这里主要描述如何使用返回值的函数

声明方式如下

`func name() string {}`

主要是在形式参数和函数体之间声明一个**变量类型**，用于确定该函数返回的是什么类型的数据

#### 2.4 数组声明

数组有两种形式

1. `array`

2. `slice`

`slcie`

用于保存同一类型数据的数组

声明内容，数组、变量类型、初值空间

示例：`cards := []string{}`

#### 2.5 数组迭代

使用，for标志、遍历指数、遍历片段、声明标志、遍历数组、遍历空间

形式：

`for i,card := range cards {}`

#### 2.6 自定义类型

声明自定义类型，形式如下：

`type name []string`

类型标识符、自定义名称、类型，

- 可以放在文件中引入，也可以放在主文件中声明

- 可以为自定义类型声明方法，并在类型被使用时提供

```go
// ----引入 
// card.go
package main

type deck []string
// main.go
package main

func main(){
    card := deck{'a','b'}
} 

// ----声明
package main

type deck []string

func main(){
    card := deck{'a','b'}
}
```

自定义类型的函数声明

声明步骤：

函数标识、类型声明、函数名称、接收参数、函数体

在类型声明这里，作用把引用该方法的对象接收过来，名称为`d`，就可以在函数体内进行使用，**名称规则由官方定义，为自定义类型的开头字母**

```go
func (d deck) print(){
    fmt.println(d)
}
```

#### 2.7 引入依赖包

实现将`[]deck -> []string -> string -> []byte`

`byte`是每个字符的ASCII码值，旨在方便计算机识别内容

`strings.Join([]string,string)`：将字符串数组转换为字符转

#### 2.8 转换数据类型

在变量之外，输入变量类型，并将需要改变类型的变量作为形式参数放入

例如：`string(bytes)    // bytes 是 []byte`

#### 2.9 随机函数

`Math`包下的`rand`对象的`Intn`方法

使用示例：`rand.Intn(n)`

`n`表示随机整数的可取上限值

#### 2.10 测试

创建一个`xx_test.go`文件，`xx`表示做测试的包名

文件内容设置如下：

```go
package main

import (
    "os"
    "testing"
)

// 测试扑克牌对象
// 1. 对象长度是否达到16
// 2. 第一张扑克牌是否为"Ace of Spades"
// 3. 最后一张扑克牌是否为"Four of Clubs"
func TestNewDeck(t *testing.T) {
    d := newDeck()

    if len(d) != 16 {
        t.Errorf("Expected deck length of 20,but got %v", len(d))
    }

    if d[0] != "Ace of Spades" {
        t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
    }

    if d[len(d)-1] != "Four of Clubs" {
        t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
    }
}


// 测试读取的文件内容
// 1. 生成扑克牌对象
// 2. 保存为文件"_decktesing"
// 3. 读取内容到 loadedDeck
// 4. loadedDeck 的长度是否为16
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
    os.Remove("_decktesing")

    deck := newDeck()
    deck.saveToFile("_decktesing")

    loadedDeck := newDeckFromFile("_decktesing")

    if len(loadedDeck) != 16 {
        t.Errorf("Expected deck length of 20,but got %v", len(loadedDeck))
    }

    os.Remove("_decktesing")
}

    }

    os.Remove("_decktesing")
}
```

---

## 第三章 数据结构

#### 3.1 声明与使用

```go
package main

import "fmt"

type person struct {
    firstName string
    lastName  string
}

func main() {
    // 声明方式一，直接声明
    //Alex := person{"Alex", "Anderson"}
    // 声明方式二，选择声明
    //Alex := person{firstName: "Alex", lastName: "Anderson"}
    // 声明方式三，结构声明
    var Alex person

    Alex.firstName = "Alex"
    Alex.lastName = "Anderson"

    fmt.Println(Alex)
}
```

#### 3.2 结构嵌套

```go
package main

import "fmt"

type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string 
    // 声明自定义类型，方式一，命名声明
    //contact   contactInfo 
    // 方式二，类型声明
    contactInfo
}

func main() {
    jim := person{
        firstName: "Jim",
        lastName:  "Bruce",
        // 名称+结构
        //contact: contactInfo{
        //    email:   "jim@gmail.com",
        //    zipCode: 95421,
        //},
        // 名称+结构
        contactInfo: contactInfo{
            email:   "jim@gmail.com",
            zipCode: 95421,
        },
    }

    //jim.print()
    jim.updateName("Bob")
    jim.print()
}

func (p person) updateName(newFirstName string) {
    p.firstName = newFirstName
}

func (p person) print() {
    fmt.Printf("%+v\n", p)
}
```

#### 3.3 数据结构更新-概念

新生成的数据结构，在使用**外部函数**更新结构内的属性值时，不会使新数据结构内的属性值发生改变，外部函数被引用时，会拷贝一个新的数据结构，并将属性值更新到拷贝数据结构的属性值上。

示例如下：

✨

```go
// 代码片段
    jim.updateName("Bob")
    jim.print()
}

func (p person) updateName(newFirstName string) {
    p.firstName = newFirstName    // 赋值发生在拷贝数据结构的属性值上
}

func (p person) print() {
    fmt.Printf("%+v\n", p)
}
```

正确示例：

```go
// 代码片段
    jimPointer := &jim // 赋予地址
    jimPointer.updateName("Bob")
    jim.print()
}

// 声明使用person的指针，并使用指针指向的数据空间进行更新
func (pointToPerson *person) updateName(newFirstName string) {
    (*pointToPerson).firstName = newFirstName
}
```

说明

- `&变量名`，表示该变量的地址指向（指针）

- `*数据类型`，表示该数据类型的地址指向

- `*地址指向`，表示该地址指向所指的数据空间

**数据类型**

基本数据类型

- int

- float

- string

- bool

- structs（自定义类型）

引用数据类型

- slice

- map

- channel

- pointer

- function

如果更新数据时，数据的数据类型是引用类型，那么将数组引用复制一份，并进行传输。

基本数据类型在进行传递时，则会将数据复制一份并进行更新

---

## 第四章 Map结构

#### 4.1 map概述

类型声明：`map[string]string`，意思表明键为字符串、值为字符串的集合

声明示例：

```go
colors := map[string]string{
        "red":   "#ff0000",
        "green": "#4bf745",
        "white": "#ffffff",
    }
```

赋值操作：`colors["red"] = "#332211"`

删除操作：`delete(colors,"red")`

Map集合与自定义结构的区别：

Map集合：

1. 键的类型相同

2. 值的类型相同

3. 可以指数迭代

4. 使用一个相关属性的集合

5. 不需要识别所有的键

6. 引用类型

自定义结构：

1. 值的类型可以不同

2. 不可以指数迭代

3. 使用不同属性代表一件东西

4. 需要识别所有的属性值

5. 基本类型

---

## 第五章 Interface(接口)
