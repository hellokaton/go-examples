package main

import "fmt"

//数组的长度是其类型的一部分，因此数组不能改变大小。
// 这看起来是一个制约，但是请不要担心； Go 提供了更加便利的方式来使用数组。
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [...]int{1,2,4}
	fmt.Println("len =", len(b), " => ", b)
}
