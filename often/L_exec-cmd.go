package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"os"
	"fmt"
)

func Exec1() {
	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个，命令参数
	cmd := exec.Command("ls", "-a", "-l")
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(opBytes))
}

func Exec2() {
	cmd := exec.Command("ls", "-a", "-l")
	// 重定向标准输出到文件
	stdout, err := os.OpenFile("stdout.log", os.O_CREATE | os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer stdout.Close()
	cmd.Stdout = stdout
	// 执行命令
	if err := cmd.Start(); err != nil {
		log.Println(err)
	}
}

func Exec3()  {
	cmd := exec.Command("sleep", "5")
	// 如果用Run，执行到该步则会阻塞等待5秒
	// err := cmd.Run()
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	// Start，上面的内容会先输出，然后这里会阻塞等待5秒
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
func main() {
	//运行系统命令
	Exec1()

	fmt.Println("========邪恶的分割线========")
	//输出重定向到文件
	Exec2()

	fmt.Println("========邪恶的分割线========")
	//	Start和Run的区别
	Exec3()
}