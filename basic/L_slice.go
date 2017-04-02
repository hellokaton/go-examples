package main

import (
	"fmt"
	"strings"
)

func slice1()  {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}

//slice 可以包含任意的类型，包括另一个 slice
func slice2()  {
	// Create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"
	printBoard(game)
}

//slice 可以重新切片，创建一个新的 slice 值指向相同的数组
func slice3()  {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])
}

//slice 由函数 make 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组
func slice4()  {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

//nil slice
func slice5()  {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

//向 slice 添加元素
func slice6()  {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}
//一个 slice 会指向一个序列的值，并且包含了长度信息
func main() {

	slice1()
	fmt.Println("======邪恶的分割线======")
	slice2()
	fmt.Println("======邪恶的分割线======")
	slice3()
	fmt.Println("======邪恶的分割线======")
	slice4()
	fmt.Println("======邪恶的分割线======")
	slice5()
	fmt.Println("======邪恶的分割线======")
	slice6()
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}