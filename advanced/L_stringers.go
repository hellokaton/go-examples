package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 类似于Java中的toString
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"爱丽丝", 18}
	z := Person{"FBI张三", 48}
	fmt.Println(a, z)
}
