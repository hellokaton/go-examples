package main

import (
	"time"
	"fmt"
)

type MyError struct {
	time time.Time
	msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v, %s",
		e.time, e.msg)
}

func run() error {
	return &MyError{time.Now(), "服务器出差错了"}
}

func main() {
	//Go 程序使用 error 值来表示错误状态
	if err := run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("没有任何异常")
	}
}