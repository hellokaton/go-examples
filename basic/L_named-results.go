package main

import "fmt"

// Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。
//返回值的名称应当具有一定的意义，可以作为文档使用。
//没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
//直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
func split(sum int) (x, y int)  {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	fmt.Println(split(17))
	var x, y = split(17)
	fmt.Println(x, "===我是可爱的分割线===", y)
}
