package main

import (
	"fmt"
	"net"
	"strings"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func startServer(port string) {
	//创建TCP监听地址
	addr, err := net.ResolveTCPAddr("tcp", port)
	checkErr(err)

	//开始监听
	listen, err := net.ListenTCP("tcp", addr)
	checkErr(err)

	fmt.Println("Server started")

	//保存客户端连接
	conns := make(map[string]net.Conn)

	//消息通道
	msg := make(chan string, 10)

	//启动协程广播消息
	go broadcast(&conns, msg)

	//启动协程处理服务端消息及命令
	go servermsg(&conns, msg)

	for {
		//接收客户端连接
		conn, err := listen.Accept()
		checkErr(err)
		//每个连接放到单独协程进行处理
		go handler(conn, msg, &conns)
	}
}

func handler(conn net.Conn, msg chan string, conns *map[string]net.Conn) {
	//初次连接发送欢迎消息
	conn.Write([]byte("Welcome to join."))
	var messages string
	data := make([]byte, 1024)

	for {
		length, err := conn.Read(data)
		if err != nil {
			conn.Close()
			break
		}
		if length > 0 {
			data[length] = 0
		}

		//解析客户端内容
		cmd := strings.Split(string(data[0:length]), "|")

		fmt.Println(string(data[0:length]))

		switch cmd[0] {
		case "hello":
			//判断是否有同名昵称并保存客户端连接
			if _, ok := (*conns)[cmd[1]]; ok {
				conn.Close()
			} else {
				(*conns)[cmd[1]] = conn
				messages = "[" + cmd[1] + "] join."
			}
		case "say":
			messages = "[" + cmd[1] + "]" + ": " + cmd[2]
		}

		//向通道发送拼装好的消息
		msg <- messages
	}

}

func broadcast(conns *map[string]net.Conn, msg chan string) {
	for {
		//从通道中接收消息
		data := <-msg

		//循环客户端连接并发送消息
		for key, value := range *conns {
			_, err := value.Write([]byte(data))
			if err != nil {
				delete(*conns, key)
			}
		}
	}
}

func servermsg(conns *map[string]net.Conn, msg chan string) {
	for {
		message := ScanLine()

		//解析命令
		cmd := strings.Split(string(message), "|")

		if len(cmd) > 1 {
			switch cmd[0] {
			case "kick":
				if _, ok := (*conns)[cmd[1]]; ok {
					//关闭对应客户端连接
					(*conns)[cmd[1]].Close()
					msg <- "[Server]: Kick [" + cmd[1] + "]"
				}
			default:
				msg <- "[Server]: " + string(message)
			}
		} else {
			msg <- "[Server]: " + string(message)
		}

	}
}