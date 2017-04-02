package main

import "fmt"

//一个结构体（ struct ）就是一个字段的集合
type Point struct {
	x, y int
}

var (
	vp1 = Point{1, 2}//类型为point
	vp2 = Point{x: 20}//y=0，被省略
	vp3 = Point{}//x=0, y=0
	mvp = &Point{1, 2}//类型为*Point
)

func main() {
	fmt.Println(Point{20, 78})

	p := Point{10, 20}
	//结构体字段使用点号来访问
	p.x = 33
	fmt.Println(p.x)

	//通过指针间接的访问是透明的
	p2 := &p
	p2.x = 99
	fmt.Println(p)

	fmt.Println(vp1, vp2, vp3, mvp)
}
