package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}


func main() {
	st := &Student {
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
	b, err := json.Marshal(st)

	var jsonstr string
	if err != nil{
		fmt.Printf("err was %v", err)
	} else {
		// 将对象解析为json字符串
		jsonstr = string(b)
		fmt.Println(jsonstr)
	}

	st2 := Student{}
	// 将字符串解析为对象
	err = json.Unmarshal([]byte(jsonstr), &st2)
	if err != nil{
		fmt.Printf("err was %v", err)
	} else {
		fmt.Println(st2)
	}
}