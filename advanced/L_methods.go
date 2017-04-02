package main

import (
	"fmt"
	"strconv"
)

//Go 没有类。然而，仍然可以在结构体类型上定义方法
type Person struct {
	name string
	age  int
}

func (p *Person) sayHello() string {
	return "I,m " + p.name + ", this year is " + strconv.Itoa(p.age)
}

type MyString string

// 你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
// 但是，不能对来自其他包的类型或基础类型定义方法。
func (str MyString) length() int {
	return len(str)
}

func main() {
	p := &Person{"jack", 23}
	fmt.Println(p.sayHello())

	fmt.Println(MyString("Hello").length())
}
