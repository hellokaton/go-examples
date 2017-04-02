package main

import (
	"os"
	"fmt"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

//类型通过实现那些方法来实现接口。 没有显式声明的必要；所以也就没有关键字“implements“
//隐式接口解藕了实现接口的包和定义接口的包：互不依赖
//因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	// os.Stdout 实现了 Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}