package main

import (
	"os"
	"fmt"
)

var userFile = "test.txt"

// 创建文件并写入字符串
func CreateFile() {

	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	fout.WriteString("Just a test!\r\n")
	fout.Write([]byte("我是一行测试文本!\r\n"))
}

func ReadFile() {
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
// 删除文件
func DeleteFile() {
	os.Remove(userFile)
}

func main() {
	CreateFile()
	ReadFile()
	DeleteFile()
}