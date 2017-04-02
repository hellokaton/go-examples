package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	//for 循环的 range 格式可以对 slice 或者 map 进行迭代循环
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("====邪恶的分割线====")

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	//可以通过赋值给 _ 来忽略序号和值
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}