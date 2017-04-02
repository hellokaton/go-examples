package main

import "fmt"

/*
Go 只有一种循环结构—— for 循环。

基本的 for 循环包含三个由分号分开的组成部分：

初始化语句：在第一次循环执行前被执行
循环条件表达式：每轮迭代开始前被求值
后置语句：每轮迭代后被执行
初始化语句一般是一个短变量声明，这里声明的变量仅在整个 for 循环语句可见。

如果条件表达式的值变为 false，那么迭代将终止。

注意：不像 C，Java，或者 Javascript 等其他语言，for 语句的三个组成部分 并不需要用括号括起来，但循环体必须用 { } 括起来。
*/

func for1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 循环初始化语句和后置语句都是可选的
func for2() {
	sum := 1
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	for1()
	for2()
}