package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// 计算字符串md5
func CalcStringMd5()  {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x\n", h.Sum(nil))
	// output: e2c569be17396eca2a2e3c11578123ed

	// 直接使用md5 ew对象的Write方式也是一样的
	h2 := md5.New()
	h2.Write([]byte("The fog is getting thicker!"))
	h2.Write([]byte("And Leon's getting laaarger!"))
	fmt.Printf("%x\n", h2.Sum(nil))
	// output: e2c569be17396eca2a2e3c11578123ed
}

// 计算文件md5
func CalcFileMd5()  {
	file, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return
	}
	fmt.Printf("%x\n", h.Sum(nil))
	// output: 43c6359298645ded23f3c2ee44acf564
}

func main() {
	CalcStringMd5()
	CalcFileMd5()
}
