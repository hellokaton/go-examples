package main

import "fmt"

// 函数可以返回任意数量的返回值。
func swap(x, y string) (string, string)  {
	return y, x
}

func main() {
	var a, b string = "hello", "golang"
	fmt.Println(swap(a, b))
}
