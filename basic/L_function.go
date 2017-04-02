package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func sum2(a, b int) int {
	return a + b
}
func main() {
	fmt.Println(sum(1, 20))
	fmt.Println(sum2(1, 20))
}
