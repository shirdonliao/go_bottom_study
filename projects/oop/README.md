

#理解golang的interface主要在于以下两点：

    interface是方法的集合
    interface是一种类型

简单示例

package main

import "fmt"

type Animal interface {
    Speak() string
}

type Cat struct{}
func (c Cat) Speak() string {
    return "cat"
}

type Dog struct{}
func (d Dog) Speak() string {
    return "dog"
}

func Test(params interface{}) {
    fmt.Println(params)
}

func main() {
    animals := []Animal{Cat{}, Dog{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }

        Test("string")
    Test(123)
    Test(true)
}

以上代码中，定义了Animal为接口，而Cat和Dog两个结构体分别实现了接口中定义的方法。当interface{}作为函数形参时，可以接受不同类型的参数。
指针与interface

如果将上述代码的：

func (c Cat) Speak() string {
    return "cat"
}

修改为:

func (c *Cat) Speak() string {
    return "cat"
}

再次运行源代码，就会出现以下错误：

cannot use Cat literal (type Cat) as type Animal in array or slice literal:
Cat does not implement Animal (Speak method has pointer receiver)

这是因为程序认为Cat并未实现Speak()方法， 而是由 *Cat 实现的。这说明结构体在实现接口方法时并不会隐式转换类型。
interface 数组

interface{} 作为函数形参和 []interface{} 作为形参有很大区别，示例如下：

package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    PrintAll(names)
}

以上代码是无法正常运行的，错误提示为：

cannot use names (type []string) as type []interface {} in argument to PrintAll

这说明对接口数组赋值前，必须多一个类型转换操作，正确代码如下：

package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}

以上是关于interface使用的基础内容。
