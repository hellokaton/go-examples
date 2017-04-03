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

/*
& 取一个变量的地址
* 取一个指针变量所指向的地址的值
所谓指针就是一个指向（存储）特定变量地址的变量
*/

func change(x *int)  {
	*x = 1024
}

func main() {
	i, j := 42, 2701
	p := &i         // 指向变量i
	fmt.Println("指针p的值:", *p) // 通过指针读取i的值
	*p = 21         // 通过指针设置i的值
	fmt.Println("变量i的值:", i)  // 查看i的新值

	p = &j         // 指向变量j
	*p = *p / 37   // 通过指针修改j的数据
	fmt.Println("变量j的值:", j) // 查看j的新值


	// 指针
	x := 10

	// 指针变量
	var x_ptr *int = &x

	var x2 = &x

	fmt.Println(x)        	// 10
	fmt.Println(x_ptr)    	// x的地址
	fmt.Println(&x)        	// x的地址
	fmt.Println(x2)        	// x的地址
	fmt.Println(*x_ptr)    	// 通过指针变量输出x的值
	fmt.Println(*x2)    	// 通过指针变量输出x的值
	fmt.Println(&x_ptr)		// 取x_ptr指针变量的地址
	fmt.Println(*&x_ptr)	// *temp =  &x_ptr	=> &x => x的地址

	num := 10
	fmt.Println(num)
	change(&num)
	fmt.Println(num)
}