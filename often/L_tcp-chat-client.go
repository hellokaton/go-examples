package main

import (
	"fmt"
	"net"
	"strings"
	"bufio"
	"os"
)

func startClient(port string) {
	//创建TCP请求地址
	addr, err := net.ResolveTCPAddr("tcp", port)
	checkErr(err)

	//连接服务端
	conn, err := net.DialTCP("tcp", nil, addr)
	checkErr(err)

	//读取欢迎信息
	data := make([]byte, 1024)
	conn.Read(data)
	fmt.Println(string(data))

	//输入昵称
	fmt.Print("Please input your nickname:")
	nickname := ScanLine()
	fmt.Println("Hello " + nickname)
	conn.Write([]byte("hello|" + nickname))

	//开启协程处理消息
	go ChatHandler(conn, nickname)

	//发送消息
	for {
		message := ScanLine()
		conn.Write([]byte("say|" + nickname + "|" + message))
	}
}

//ScanLine 读取整行
func ScanLine() string {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	return strings.Replace(input, "\n", "", -1)
}

func ChatHandler(conn net.Conn, nickname string) {
	for {
		data := make([]byte, 1024)
		//读取消息
		_, err := conn.Read(data)
		checkErr(err)
		//屏蔽自身发送的聊天内容
		if strings.Contains(string(data), "["+nickname+"]: ") == false {
			fmt.Println(string(data))
		}
	}
}