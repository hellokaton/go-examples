package main

import "fmt"

/*
Go 具有指针。 指针保存了变量的内存地址。

类型 *T 是指向类型 T 的值的指针。其零值是 nil 。

var p *int
& 符号会生成一个指向其作用对象的指针。

i := 42
p = &i
* 符号表示指针指向的底层的值。

fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的“间接引用”或“非直接引用”。

与 C 不同，Go 没有指针运算。
*/
func main() {
	i, j := 42, 2701
	p := &i         // 指向变量i
	fmt.Println("指针p的值:", *p) // 通过指针读取i的值
	*p = 21         // 通过指针设置i的值
	fmt.Println("变量i的值:", i)  // 查看i的新值

	p = &j         // 指向变量j
	*p = *p / 37   // 通过指针修改j的数据
	fmt.Println("变量j的值:", j) // 查看j的新值
}