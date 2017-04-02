package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

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

//基于此可以省略分号：C 的 while 在 Go 中叫做 for
func while1() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// 就像 for 循环一样，Go 的 if 语句也不要求用 ( ) 将条件括起来，同时， { } 还是必须有的。
func if1(x int) {
	if x > 0 {
		fmt.Println("x > 0")
	} else {
		fmt.Println("x <= 0")
	}
}

//跟 for 一样， if 语句可以在条件之前执行一个简单语句
func if2(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

func switch1() {
	fmt.Print("当前操作系统是：")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
//switch 的条件从上到下的执行，当匹配成功的时候停止
func switch2() {
	fmt.Println("Q:今天是周六吗?")
	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0:
			fmt.Println("A:今天是周六")
		case today + 1:
			fmt.Println("A:明天是周六.")
		case today + 2:
			fmt.Println("A:还有2天.")
		default:
			fmt.Println("A:早着呢.")
	}
}

//没有条件的 switch 同 switch true 一样
func switch3()  {
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("早上好!")
		case t.Hour() < 17:
			fmt.Println("下午好.")
		default:
			fmt.Println("晚上好.")
	}
}
func main() {

	for1()
	for2()

	while1()

	if1(20)
	fmt.Println(if2(3, 3, 20))

	switch1()
	switch2()
	switch3()
}